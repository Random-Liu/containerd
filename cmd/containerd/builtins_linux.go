package main

import (
	_ "github.com/containerd/containerd/linux"
	_ "github.com/containerd/containerd/metrics/cgroups"
	_ "github.com/containerd/containerd/snapshot/naive"
	_ "github.com/containerd/containerd/snapshot/overlay"
)
