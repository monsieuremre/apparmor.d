// apparmor.d - Full set of apparmor profiles
// Copyright (C) 2021-2023 Alexandre Pujol <alexandre@pujol.io>
// SPDX-License-Identifier: GPL-2.0-only

package aa

var (
	// Include
	include1      = &Include{IsMagic: true, Path: "abstraction/base"}
	include2      = &Include{IsMagic: false, Path: "abstraction/base"}
	include3      = &Include{IfExists: true, IsMagic: true, Path: "abstraction/base"}
	includeLocal1 = &Include{IfExists: true, IsMagic: true, Path: "local/foo"}

	// Rlimit
	rlimit1 = &Rlimit{Key: "nproc", Op: "<=", Value: "200"}
	rlimit2 = &Rlimit{Key: "cpu", Op: "<=", Value: "2"}
	rlimit3 = &Rlimit{Key: "nproc", Op: "<", Value: "2"}

	// Capability
	capability1Log = map[string]string{
		"apparmor":   "ALLOWED",
		"class":      "cap",
		"operation":  "capable",
		"capname":    "net_admin",
		"capability": "12",
		"profile":    "pkexec",
		"comm":       "pkexec",
	}
	capability1 = &Capability{Name: "net_admin"}
	capability2 = &Capability{Name: "sys_ptrace"}

	// Network
	network1Log = map[string]string{
		"apparmor":       "ALLOWED",
		"class":          "net",
		"operation":      "create",
		"family":         "netlink",
		"profile":        "sddm-greeter",
		"sock_type":      "raw",
		"protocol":       "15",
		"requested_mask": "create",
		"denied_mask":    "create",
		"comm":           "sddm-greeter",
	}
	network1 = &Network{Domain: "netlink", Type: "raw", Protocol: "15"}
	network2 = &Network{Domain: "inet", Type: "dgram"}

	// Mount
	mount1Log = map[string]string{
		"apparmor":  "ALLOWED",
		"class":     "mount",
		"operation": "mount",
		"info":      "failed perms check",
		"error":     "-13",
		"profile":   "dockerd",
		"name":      "/var/lib/docker/overlay2/opaque-bug-check1209538631/merged/",
		"comm":      "dockerd",
		"fstype":    "overlay",
		"srcname":   "overlay",
	}
	mount2Log = map[string]string{
		"apparmor":  "ALLOWED",
		"class":     "mount",
		"operation": "mount",
		"info":      "failed perms check",
		"error":     "-13",
		"profile":   "dockerd",
		"name":      "/var/lib/docker/overlay2/metacopy-check906831159/merged/",
		"comm":      "dockerd",
		"fstype":    "overlay",
		"srcname":   "overlay",
	}
	mount1 = &Mount{
		Qualifier:       Qualifier{Comment: "failed perms check"},
		MountConditions: MountConditions{FsType: "overlay", Options: []string{}},
		Source:          "overlay",
		MountPoint:      "/var/lib/docker/overlay2/opaque-bug-check1209538631/merged/",
	}
	mount2 = &Mount{
		Qualifier:       Qualifier{Comment: "failed perms check"},
		MountConditions: MountConditions{FsType: "overlay", Options: []string{}},
		Source:          "overlay",
		MountPoint:      "/var/lib/docker/overlay2/metacopy-check906831159/merged/",
	}

	// Change Profile
	changeprofile1 = &ChangeProfile{ExecMode: "Px", Exec: "/bin/bash", ProfileName: "brwap//default"}
	changeprofile2 = &ChangeProfile{ExecMode: "Px", Exec: "/bin/bash", ProfileName: "brwap"}
	changeprofile3 = &ChangeProfile{ExecMode: "safe", Exec: "/bin/foo", ProfileName: "brwap//default"}

	// Signal
	signal1Log = map[string]string{
		"apparmor":       "ALLOWED",
		"class":          "signal",
		"profile":        "firefox",
		"operation":      "signal",
		"comm":           "49504320492F4F20506172656E74",
		"requested_mask": "receive",
		"denied_mask":    "receive",
		"signal":         "kill",
		"peer":           "firefox//&firejail-default",
	}
	signal1 = &Signal{
		Access: "receive",
		Set:    "kill",
		Peer:   "firefox//&firejail-default",
	}
	signal2 = &Signal{
		Access: "receive",
		Set:    "up",
		Peer:   "firefox//&firejail-default",
	}

	// Ptrace
	ptrace1Log = map[string]string{
		"apparmor":       "ALLOWED",
		"class":          "ptrace",
		"profile":        "xdg-document-portal",
		"operation":      "ptrace",
		"comm":           "pool-/usr/lib/x",
		"requested_mask": "read",
		"denied_mask":    "read",
		"peer":           "nautilus",
	}
	ptrace2Log = map[string]string{
		"apparmor":       "DENIED",
		"class":          "ptrace",
		"operation":      "ptrace",
		"comm":           "systemd-journal",
		"requested_mask": "readby",
		"denied_mask":    "readby",
		"peer":           "systemd-journald",
	}
	ptrace1 = &Ptrace{Access: "read", Peer: "nautilus"}
	ptrace2 = &Ptrace{Access: "readby", Peer: "systemd-journald"}

	// Unix
	unix1Log = map[string]string{
		"apparmor":       "ALLOWED",
		"class":          "unix",
		"family":         "unix",
		"operation":      "file_perm",
		"profile":        "gsettings",
		"comm":           "dbus-daemon",
		"requested_mask": "send receive",
		"addr":           "none",
		"peer_addr":      "@/tmp/dbus-AaKMpxzC4k",
		"peer":           "dbus-daemon",
		"denied_mask":    "send receive",
		"sock_type":      "stream",
		"protocol":       "0",
	}
	unix1 = &Unix{
		Access:   "send receive",
		Type:     "stream",
		Protocol: "0",
		Address:  "none",
		Peer:     "dbus-daemon",
		PeerAddr: "@/tmp/dbus-AaKMpxzC4k",
	}
	unix2 = &Unix{
		Qualifier: Qualifier{FileInherit: true},
		Access:    "receive",
		Type:      "stream",
	}

	// Dbus
	dbus1Log = map[string]string{
		"apparmor":   "ALLOWED",
		"operation":  "dbus_method_call",
		"bus":        "session",
		"path":       "/org/gtk/vfs/metadata",
		"interface":  "org.gtk.vfs.Metadata",
		"member":     "Remove",
		"name":       ":1.15",
		"mask":       "receive",
		"label":      "gvfsd-metadata",
		"peer_pid":   "3888",
		"peer_label": "tracker-extract",
	}
	dbus2Log = map[string]string{
		"apparmor":  "ALLOWED",
		"operation": "dbus_bind",
		"bus":       "session",
		"name":      "org.gnome.evolution.dataserver.Sources5",
		"mask":      "bind",
		"pid":       "3442",
		"label":     "evolution-source-registry",
	}
	dbus1 = &Dbus{
		Access:    "receive",
		Bus:       "session",
		Name:      ":1.15",
		Path:      "/org/gtk/vfs/metadata",
		Interface: "org.gtk.vfs.Metadata",
		Member:    "Remove",
		Label:     "tracker-extract",
	}
	dbus2 = &Dbus{
		Access: "bind",
		Bus:    "session",
		Name:   "org.gnome.evolution.dataserver.Sources5",
	}
	dbus3 = &Dbus{
		Access: "bind",
		Bus:    "session",
		Name:   "org.gnome.evolution.dataserver",
	}

	// File
	file1Log = map[string]string{
		"apparmor":       "ALLOWED",
		"operation":      "open",
		"class":          "file",
		"profile":        "cupsd",
		"name":           "/usr/share/poppler/cMap/Identity-H",
		"comm":           "gs",
		"requested_mask": "r",
		"denied_mask":    "r",
		"fsuid":          "209",
		"FSUID":          "cups",
		"ouid":           "0",
		"OUID":           "root",
	}
	file2Log = map[string]string{
		"apparmor":       "ALLOWED",
		"operation":      "open",
		"class":          "file",
		"profile":        "gsd-print-notifications",
		"name":           "@{PROC}/4163/cgroup",
		"comm":           "gsd-print-notif",
		"requested_mask": "r",
		"denied_mask":    "r",
		"fsuid":          "1000",
		"FSUID":          "user",
		"ouid":           "1000",
		"OUID":           "user",
		"error":          "-1",
	}
	file1 = &File{Path: "/usr/share/poppler/cMap/Identity-H", Access: "r"}
	file2 = &File{
		Qualifier: Qualifier{Owner: true, NoNewPrivs: true},
		Path:      "@{PROC}/4163/cgroup",
		Access:    "r",
	}
)
