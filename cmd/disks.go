package main

import (
	"fmt"
	"sort"

	"github.com/jedib0t/go-pretty/table"
	"github.com/rchamarthy/storix"
	"github.com/spf13/cobra"
	"machinerun.io/disko"
	"machinerun.io/disko/linux"
)

const NA = "--"

func NewShowDisks() *cobra.Command {
	handler := func(c *cobra.Command, args []string) error {
		sys := linux.System()
		disks, err := sys.ScanAllDisks(func(disko.Disk) bool { return true })
		if err != nil {
			return err
		}

		showDisksTable(disks)

		return nil
	}

	return &cobra.Command{
		Use:   "disks",
		Short: "Show disk information",
		RunE:  handler,
	}
}

// showDisks displays information about all the disks found in the system.
func showDisksTable(disks disko.DiskSet) {
	oDisks := sortDisks(disks)

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Name", "Media", "Attachment", "Size", "FreeSpace"})
	tw.Style().Options.DrawBorder = false

	fs := func(f disko.Disk) string {
		var avial uint64
		fs := f.FreeSpaces()
		for _, f := range fs {
			avial += f.Size()
		}
		return sizeStr(avial)
	}

	for _, d := range oDisks {
		s := storix.ByteSize(d.Size).String()
		tw.AppendRow(table.Row{d.Name, d.Type.String(), d.Attachment.String(),
			s, fs(d)})
	}

	fmt.Println(tw.Render())
}

func sizeStr(size uint64) string {
	return storix.ByteSize(size).String()
}

func sortDisks(disks disko.DiskSet) []disko.Disk {
	oDisks := []string{}
	for idx := range disks {
		oDisks = append(oDisks, disks[idx].Name)
	}
	sort.Strings(oDisks)

	sortedDisks := []disko.Disk{}
	for _, n := range oDisks {
		sortedDisks = append(sortedDisks, disks[n])
	}

	return sortedDisks
}
