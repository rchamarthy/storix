package main

import (
	"fmt"
	"sort"

	"github.com/jedib0t/go-pretty/table"
	"github.com/spf13/cobra"
	"machinerun.io/disko"
	"machinerun.io/disko/linux"
	"machinerun.io/disko/partid"
)

func NewShowPartitions() *cobra.Command {
	handler := func(c *cobra.Command, args []string) error {
		sys := linux.System()
		disks, err := sys.ScanAllDisks(func(disko.Disk) bool { return true })
		if err != nil {
			return err
		}

		showPartitionsTable(disks)

		return nil
	}

	return &cobra.Command{
		Use:   "partitions",
		Short: "Show disk partition information",
		RunE:  handler,
	}
}

func showPartitionsTable(disks disko.DiskSet) {
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Label", "Name", "Media", "Attachment", "Size", "Disk Size", "Type"})
	tw.Style().Options.DrawBorder = false

	partitions := []struct {
		d disko.Disk
		p uint
	}{}

	// Collect all the partitions
	for _, d := range disks {
		for p := range d.Partitions {
			partitions = append(partitions, struct {
				d disko.Disk
				p uint
			}{d, p})
		}
	}

	// Sort them by label
	sort.Slice(partitions, func(i, j int) bool {
		p1 := partitions[i].d.Partitions[partitions[i].p]
		p2 := partitions[j].d.Partitions[partitions[j].p]

		return p1.Name < p2.Name
	})

	for _, p := range partitions {
		partitionDetails(tw, p.d, p.p)
	}

	fmt.Println(tw.Render())
}

func partitionDetails(tw table.Writer, d disko.Disk, pNum uint) {
	p := d.Partitions[pNum]
	pt := partid.Text[p.Type]
	if pt == "" {
		pt = NA
	}

	n := NA
	if p.Name != "" {
		n = p.Name
	}

	pName := fmt.Sprintf("%s%d", d.Name, p.Number)
	row := table.Row{n, pName, d.Type.String(), d.Attachment.String(), sizeStr(p.Size()), sizeStr(d.Size), pt}
	tw.AppendRow(row)
}
