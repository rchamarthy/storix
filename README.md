# stroix - storage tool for linux

storix is designed to perform automation of storage management tasks for linux
appliances. This tool can be used to create storage volumes using LVM to store
application specific data.

__Functionality:__

* Storage profiles - Define multiple storage profiles and use them based on the
actual storage capacity presented by the underlying storage devices.
* Disk Selection - Select specific disks to be used for storage and allocate
volumes __only__ on those disks.
* Volume Groups - Define volume groups based for a disk pool.
* Volumes - Allocate volumes, create the specified file system and mount them
at a defined mount point in the root file system.

## Storage Profile

A stroix storage profile document defines the disks, volume groups and volumes
for a given storage profile. There can be multiple storage profiles defined but
only one storage profile can be used at a time. Profile selection is usually
based on the appliance type or product type and is done external to storix tool.
Storix expects the profile file to be present at `/etc/storix/storage.yaml`.

```yaml
storix:
  disks:
    - disk-pool: fast-disks
      match:
        - type: ssd
          min: 100GiB
          max: 1TiB
          attachement: RAID
        - type: ssd
          min: 100GiB
          max: 1TiB
          attachement: PCIE
    - disk-pool: slow-disks
      match:
        - type: hdd
          min: 1TiB
          max: 3TiB
          attachement: SATA
  volume-groups:
    - name: fast0
      label: fast
      disk-pool: fast-disks
    - name: slow0
      label: slow
      disk-pool: slow-disks
  volumes:
    - name: logs
      group: fast
      size: 100GiB
      fs: ext4
      mount: /var/log
    - name: techsupport
      group: slow
      size: 40GiB
      fs: ext4
      mount: /var/techsupport
```

The profile can be changed in a compatible way like adding new disks, new
volume groups and new volumes. Storix will automatically detect __only__ the
operations that need to done and perform those. In other words, storix can do
create, update for compatible changes and retrieve operations. Deletion of
volumes, volume groups and disks is not supported. A blanket storage wipe
command is provided to wipe all the storage and start fresh. This is good
enough for an appliance use case where the state of the application running
on the appliance can be wiped all at once and started fresh.

> :warning: __Note:__ Only blank disks are considered for disk selection, if
there is a partition table present on the disk, it will be ignored. Users can
either `wipe-disk` or use the explicit `add-disk` command to manually add
a disk with existing partition table to the storage pool.

## Commands

1. Select all the storix disks that have the storix partition and wipe the
partition table and make the disk a blank disk. Naturally all volumes and
volume groups that are on these disks will be wiped.

```bash
storix wipe-disk <disk>
```

1. Add a disk or a partition to the storage pool. This command is useful when
a disk is not or there are pre-created partitions that can be used for storage.

```bash
storix add-device <device> <disk-pool>
```

1. Allocate the disks, volume groups and volumes based on the profile.

```bash
storix allocate
```

1. Clean only the partitions allocated by the current storix profile.

```bash
storix clean
```

1. Show commands to display storage information.

```bash
storix show disks [--all] [--allocated] [--unallocated]
storix show vg [--label <label>]
storix show volumes [--group <label>]
storix show status
```
