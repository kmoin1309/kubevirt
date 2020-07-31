// Code generated by swagger-doc. DO NOT EDIT.

package v1

func (HostDisk) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "Represents a disk created on the cluster level\n\n+k8s:openapi-gen=true",
		"path":     "The path to HostDisk image located on the cluster",
		"type":     "Contains information if disk.img exists or should be created\nallowed options are 'Disk' and 'DiskOrCreate'",
		"capacity": "Capacity of the sparse disk\n+optional",
		"shared":   "Shared indicate whether the path is shared between nodes",
	}
}

func (ConfigMapVolumeSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":            "ConfigMapVolumeSource adapts a ConfigMap into a volume.\nMore info: https://kubernetes.io/docs/concepts/storage/volumes/#configmap\n\n+k8s:openapi-gen=true",
		"optional":    "Specify whether the ConfigMap or it's keys must be defined\n+optional",
		"volumeLabel": "The volume label of the resulting disk inside the VMI.\nDifferent bootstrapping mechanisms require different values.\nTypical values are \"cidata\" (cloud-init), \"config-2\" (cloud-init) or \"OEMDRV\" (kickstart).\n+optional",
	}
}

func (SecretVolumeSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":            "SecretVolumeSource adapts a Secret into a volume.\n\n+k8s:openapi-gen=true",
		"secretName":  "Name of the secret in the pod's namespace to use.\nMore info: https://kubernetes.io/docs/concepts/storage/volumes#secret\n+optional",
		"optional":    "Specify whether the Secret or it's keys must be defined\n+optional",
		"volumeLabel": "The volume label of the resulting disk inside the VMI.\nDifferent bootstrapping mechanisms require different values.\nTypical values are \"cidata\" (cloud-init), \"config-2\" (cloud-init) or \"OEMDRV\" (kickstart).\n+optional",
	}
}

func (DownwardAPIVolumeSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":            "DownwardAPIVolumeSource represents a volume containing downward API info.\n\n+k8s:openapi-gen=true",
		"fields":      "Fields is a list of downward API volume file\n+optional",
		"volumeLabel": "The volume label of the resulting disk inside the VMI.\nDifferent bootstrapping mechanisms require different values.\nTypical values are \"cidata\" (cloud-init), \"config-2\" (cloud-init) or \"OEMDRV\" (kickstart).\n+optional",
	}
}

func (ServiceAccountVolumeSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                   "ServiceAccountVolumeSource adapts a ServiceAccount into a volume.\n\n+k8s:openapi-gen=true",
		"serviceAccountName": "Name of the service account in the pod's namespace to use.\nMore info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/",
	}
}

func (CloudInitNoCloudSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                     "Represents a cloud-init nocloud user data source.\nMore info: http://cloudinit.readthedocs.io/en/latest/topics/datasources/nocloud.html\n\n+k8s:openapi-gen=true",
		"secretRef":            "UserDataSecretRef references a k8s secret that contains NoCloud userdata.\n+ optional",
		"userDataBase64":       "UserDataBase64 contains NoCloud cloud-init userdata as a base64 encoded string.\n+ optional",
		"userData":             "UserData contains NoCloud inline cloud-init userdata.\n+ optional",
		"networkDataSecretRef": "NetworkDataSecretRef references a k8s secret that contains NoCloud networkdata.\n+ optional",
		"networkDataBase64":    "NetworkDataBase64 contains NoCloud cloud-init networkdata as a base64 encoded string.\n+ optional",
		"networkData":          "NetworkData contains NoCloud inline cloud-init networkdata.\n+ optional",
	}
}

func (CloudInitConfigDriveSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                     "Represents a cloud-init config drive user data source.\nMore info: https://cloudinit.readthedocs.io/en/latest/topics/datasources/configdrive.html\n\n+k8s:openapi-gen=true",
		"secretRef":            "UserDataSecretRef references a k8s secret that contains config drive userdata.\n+ optional",
		"userDataBase64":       "UserDataBase64 contains config drive cloud-init userdata as a base64 encoded string.\n+ optional",
		"userData":             "UserData contains config drive inline cloud-init userdata.\n+ optional",
		"networkDataSecretRef": "NetworkDataSecretRef references a k8s secret that contains config drive networkdata.\n+ optional",
		"networkDataBase64":    "NetworkDataBase64 contains config drive cloud-init networkdata as a base64 encoded string.\n+ optional",
		"networkData":          "NetworkData contains config drive inline cloud-init networkdata.\n+ optional",
	}
}

func (DomainSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                "+k8s:openapi-gen=true",
		"resources":       "Resources describes the Compute Resources required by this vmi.",
		"cpu":             "CPU allow specified the detailed CPU topology inside the vmi.\n+optional",
		"memory":          "Memory allow specifying the VMI memory features.\n+optional",
		"machine":         "Machine type.\n+optional",
		"firmware":        "Firmware.\n+optional",
		"clock":           "Clock sets the clock and timers of the vmi.\n+optional",
		"features":        "Features like acpi, apic, hyperv, smm.\n+optional",
		"devices":         "Devices allows adding disks, network interfaces, and others",
		"ioThreadsPolicy": "Controls whether or not disks will share IOThreads.\nOmitting IOThreadsPolicy disables use of IOThreads.\nOne of: shared, auto\n+optional",
		"chassis":         "Chassis specifies the chassis info passed to the domain.\n+optional",
	}
}

func (Chassis) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "Chassis specifies the chassis info passed to the domain.\n\n+k8s:openapi-gen=true",
	}
}

func (Bootloader) SwaggerDoc() map[string]string {
	return map[string]string{
		"":     "Represents the firmware blob used to assist in the domain creation process.\nUsed for setting the QEMU BIOS file path for the libvirt domain.\n\n+k8s:openapi-gen=true",
		"bios": "If set (default), BIOS will be used.\n+optional",
		"efi":  "If set, EFI will be used instead of BIOS.\n+optional",
	}
}

func (BIOS) SwaggerDoc() map[string]string {
	return map[string]string{
		"":          "If set (default), BIOS will be used.\n\n+k8s:openapi-gen=true",
		"useSerial": "If set, the BIOS output will be transmitted over serial\n+optional",
	}
}

func (EFI) SwaggerDoc() map[string]string {
	return map[string]string{
		"":           "If set, EFI will be used instead of BIOS.\n\n+k8s:openapi-gen=true",
		"secureBoot": "If set, SecureBoot will be enabled and the OVMF roms will be swapped for\nSecureBoot-enabled ones.\nRequires SMM to be enabled.\nDefaults to true\n+optional",
	}
}

func (ResourceRequirements) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                        "+k8s:openapi-gen=true",
		"requests":                "Requests is a description of the initial vmi resources.\nValid resource keys are \"memory\" and \"cpu\".\n+optional",
		"limits":                  "Limits describes the maximum amount of compute resources allowed.\nValid resource keys are \"memory\" and \"cpu\".\n+optional",
		"overcommitGuestOverhead": "Don't ask the scheduler to take the guest-management overhead into account. Instead\nput the overhead only into the container's memory limit. This can lead to crashes if\nall memory is in use on a node. Defaults to false.",
	}
}

func (CPU) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                      "CPU allows specifying the CPU topology.\n\n+k8s:openapi-gen=true",
		"cores":                 "Cores specifies the number of cores inside the vmi.\nMust be a value greater or equal 1.",
		"sockets":               "Sockets specifies the number of sockets inside the vmi.\nMust be a value greater or equal 1.",
		"threads":               "Threads specifies the number of threads inside the vmi.\nMust be a value greater or equal 1.",
		"model":                 "Model specifies the CPU model inside the VMI.\nList of available models https://github.com/libvirt/libvirt/tree/master/src/cpu_map.\nIt is possible to specify special cases like \"host-passthrough\" to get the same CPU as the node\nand \"host-model\" to get CPU closest to the node one.\nDefaults to host-model.\n+optional",
		"features":              "Features specifies the CPU features list inside the VMI.\n+optional",
		"dedicatedCpuPlacement": "DedicatedCPUPlacement requests the scheduler to place the VirtualMachineInstance on a node\nwith enough dedicated pCPUs and pin the vCPUs to it.\n+optional",
		"isolateEmulatorThread": "IsolateEmulatorThread requests one more dedicated pCPU to be allocated for the VMI to place\nthe emulator thread on it.\n+optional",
	}
}

func (CPUFeature) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "CPUFeature allows specifying a CPU feature.\n\n+k8s:openapi-gen=true",
		"name":   "Name of the CPU feature",
		"policy": "Policy is the CPU feature attribute which can have the following attributes:\nforce    - The virtual CPU will claim the feature is supported regardless of it being supported by host CPU.\nrequire  - Guest creation will fail unless the feature is supported by the host CPU or the hypervisor is able to emulate it.\noptional - The feature will be supported by virtual CPU if and only if it is supported by host CPU.\ndisable  - The feature will not be supported by virtual CPU.\nforbid   - Guest creation will fail if the feature is supported by host CPU.\nDefaults to require\n+optional",
	}
}

func (Memory) SwaggerDoc() map[string]string {
	return map[string]string{
		"":          "Memory allows specifying the VirtualMachineInstance memory features.\n\n+k8s:openapi-gen=true",
		"hugepages": "Hugepages allow to use hugepages for the VirtualMachineInstance instead of regular memory.\n+optional",
		"guest":     "Guest allows to specifying the amount of memory which is visible inside the Guest OS.\nThe Guest must lie between Requests and Limits from the resources section.\nDefaults to the requested memory in the resources section if not specified.\n+ optional",
	}
}

func (Hugepages) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "Hugepages allow to use hugepages for the VirtualMachineInstance instead of regular memory.\n\n+k8s:openapi-gen=true",
		"pageSize": "PageSize specifies the hugepage size, for x86_64 architecture valid values are 1Gi and 2Mi.",
	}
}

func (Machine) SwaggerDoc() map[string]string {
	return map[string]string{
		"":     "+k8s:openapi-gen=true",
		"type": "QEMU machine type is the actual chipset of the VirtualMachineInstance.",
	}
}

func (Firmware) SwaggerDoc() map[string]string {
	return map[string]string{
		"":           "+k8s:openapi-gen=true",
		"uuid":       "UUID reported by the vmi bios.\nDefaults to a random generated uid.",
		"bootloader": "Settings to control the bootloader that is used.\n+optional",
		"serial":     "The system-serial-number in SMBIOS",
	}
}

func (Devices) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                           "+k8s:openapi-gen=true",
		"disks":                      "Disks describes disks, cdroms, floppy and luns which are connected to the vmi.",
		"watchdog":                   "Watchdog describes a watchdog device which can be added to the vmi.",
		"interfaces":                 "Interfaces describe network interfaces which are added to the vmi.",
		"inputs":                     "Inputs describe input devices",
		"autoattachPodInterface":     "Whether to attach a pod network interface. Defaults to true.",
		"autoattachGraphicsDevice":   "Whether to attach the default graphics device or not.\nVNC will not be available if set to false. Defaults to true.",
		"autoattachSerialConsole":    "Whether to attach the default serial console or not.\nSerial console access will not be available if set to false. Defaults to true.",
		"autoattachMemBalloon":       "Whether to attach the Memory balloon device with default period.\nPeriod can be adjusted in virt-config.\nDefaults to true.\n+optional",
		"rng":                        "Whether to have random number generator from host\n+optional",
		"blockMultiQueue":            "Whether or not to enable virtio multi-queue for block devices\n+optional",
		"networkInterfaceMultiqueue": "If specified, virtual network interfaces configured with a virtio bus will also enable the vhost multiqueue feature for network devices. The number of queues created depends on additional factors of the VirtualMachineInstance, like the number of guest CPUs.\n+optional",
		"gpus":                       "Whether to attach a GPU device to the vmi.\n+optional\n+listType=atomic",
		"filesystems":                "Filesystems describes filesystem which is connected to the vmi.\n+optional\n+listType=atomic",
		"hostDevices":                "Whether to attach a host device to the vmi.\n+optional\n+listType=atomic",
	}
}

func (Input) SwaggerDoc() map[string]string {
	return map[string]string{
		"":     "+k8s:openapi-gen=true",
		"bus":  "Bus indicates the bus of input device to emulate.\nSupported values: virtio, usb.",
		"type": "Type indicated the type of input device.\nSupported values: tablet.",
		"name": "Name is the device name",
	}
}

func (Filesystem) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "+k8s:openapi-gen=true",
		"name":     "Name is the device name",
		"virtiofs": "Virtiofs is supported",
	}
}

func (FilesystemVirtiofs) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "+k8s:openapi-gen=true",
	}
}

func (GPU) SwaggerDoc() map[string]string {
	return map[string]string{
		"":     "+k8s:openapi-gen=true",
		"name": "Name of the GPU device as exposed by a device plugin",
	}
}

func (HostDevice) SwaggerDoc() map[string]string {
	return map[string]string{
		"":           "+k8s:openapi-gen=true",
		"deviceName": "DeviceName is the resource name of the host device exposed by a device plugin",
	}
}

func (Disk) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                  "+k8s:openapi-gen=true",
		"name":              "Name is the device name",
		"bootOrder":         "BootOrder is an integer value > 0, used to determine ordering of boot devices.\nLower values take precedence.\nEach disk or interface that has a boot order must have a unique value.\nDisks without a boot order are not tried if a disk with a boot order exists.\n+optional",
		"serial":            "Serial provides the ability to specify a serial number for the disk device.\n+optional",
		"dedicatedIOThread": "dedicatedIOThread indicates this disk should have an exclusive IO Thread.\nEnabling this implies useIOThreads = true.\nDefaults to false.\n+optional",
		"cache":             "Cache specifies which kvm disk cache mode should be used.\n+optional",
		"io":                "IO specifies which QEMU disk IO mode should be used.\nSupported values are: native, default, threads.\n+optional",
		"tag":               "If specified, disk address and its tag will be provided to the guest via config drive metadata\n+optional",
	}
}

func (DiskDevice) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "Represents the target of a volume to mount.\nOnly one of its members may be specified.\n\n+k8s:openapi-gen=true",
		"disk":   "Attach a volume as a disk to the vmi.",
		"lun":    "Attach a volume as a LUN to the vmi.",
		"floppy": "Attach a volume as a floppy to the vmi.",
		"cdrom":  "Attach a volume as a cdrom to the vmi.",
	}
}

func (DiskTarget) SwaggerDoc() map[string]string {
	return map[string]string{
		"":           "+k8s:openapi-gen=true",
		"bus":        "Bus indicates the type of disk device to emulate.\nsupported values: virtio, sata, scsi.",
		"readonly":   "ReadOnly.\nDefaults to false.",
		"pciAddress": "If specified, the virtual disk will be placed on the guests pci address with the specifed PCI address. For example: 0000:81:01.10\n+optional",
	}
}

func (LunTarget) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "+k8s:openapi-gen=true",
		"bus":      "Bus indicates the type of disk device to emulate.\nsupported values: virtio, sata, scsi.",
		"readonly": "ReadOnly.\nDefaults to false.",
	}
}

func (FloppyTarget) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "+k8s:openapi-gen=true",
		"readonly": "ReadOnly.\nDefaults to false.",
		"tray":     "Tray indicates if the tray of the device is open or closed.\nAllowed values are \"open\" and \"closed\".\nDefaults to closed.\n+optional",
	}
}

func (CDRomTarget) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "+k8s:openapi-gen=true",
		"bus":      "Bus indicates the type of disk device to emulate.\nsupported values: virtio, sata, scsi.",
		"readonly": "ReadOnly.\nDefaults to true.",
		"tray":     "Tray indicates if the tray of the device is open or closed.\nAllowed values are \"open\" and \"closed\".\nDefaults to closed.\n+optional",
	}
}

func (Volume) SwaggerDoc() map[string]string {
	return map[string]string{
		"":     "Volume represents a named volume in a vmi.\n\n+k8s:openapi-gen=true",
		"name": "Volume's name.\nMust be a DNS_LABEL and unique within the vmi.\nMore info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
	}
}

func (VolumeSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                      "Represents the source of a volume to mount.\nOnly one of its members may be specified.\n\n+k8s:openapi-gen=true",
		"hostDisk":              "HostDisk represents a disk created on the cluster level\n+optional",
		"persistentVolumeClaim": "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace.\nDirectly attached to the vmi via qemu.\nMore info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims\n+optional",
		"cloudInitNoCloud":      "CloudInitNoCloud represents a cloud-init NoCloud user-data source.\nThe NoCloud data will be added as a disk to the vmi. A proper cloud-init installation is required inside the guest.\nMore info: http://cloudinit.readthedocs.io/en/latest/topics/datasources/nocloud.html\n+optional",
		"cloudInitConfigDrive":  "CloudInitConfigDrive represents a cloud-init Config Drive user-data source.\nThe Config Drive data will be added as a disk to the vmi. A proper cloud-init installation is required inside the guest.\nMore info: https://cloudinit.readthedocs.io/en/latest/topics/datasources/configdrive.html\n+optional",
		"containerDisk":         "ContainerDisk references a docker image, embedding a qcow or raw disk.\nMore info: https://kubevirt.gitbooks.io/user-guide/registry-disk.html\n+optional",
		"ephemeral":             "Ephemeral is a special volume source that \"wraps\" specified source and provides copy-on-write image on top of it.\n+optional",
		"emptyDisk":             "EmptyDisk represents a temporary disk which shares the vmis lifecycle.\nMore info: https://kubevirt.gitbooks.io/user-guide/disks-and-volumes.html\n+optional",
		"dataVolume":            "DataVolume represents the dynamic creation a PVC for this volume as well as\nthe process of populating that PVC with a disk image.\n+optional",
		"configMap":             "ConfigMapSource represents a reference to a ConfigMap in the same namespace.\nMore info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-pod-configmap/\n+optional",
		"secret":                "SecretVolumeSource represents a reference to a secret data in the same namespace.\nMore info: https://kubernetes.io/docs/concepts/configuration/secret/\n+optional",
		"downwardAPI":           "DownwardAPI represents downward API about the pod that should populate this volume\n+optional",
		"serviceAccount":        "ServiceAccountVolumeSource represents a reference to a service account.\nThere can only be one volume of this type!\nMore info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/\n+optional",
	}
}

func (DataVolumeSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":     "+k8s:openapi-gen=true",
		"name": "Name represents the name of the DataVolume in the same namespace",
	}
}

func (EphemeralVolumeSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                      "+k8s:openapi-gen=true",
		"persistentVolumeClaim": "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace.\nDirectly attached to the vmi via qemu.\nMore info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims\n+optional",
	}
}

func (EmptyDiskSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "EmptyDisk represents a temporary disk which shares the vmis lifecycle.\n\n+k8s:openapi-gen=true",
		"capacity": "Capacity of the sparse disk.",
	}
}

func (ContainerDiskSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                "Represents a docker image with an embedded disk.\n\n+k8s:openapi-gen=true",
		"image":           "Image is the name of the image with the embedded disk.",
		"imagePullSecret": "ImagePullSecret is the name of the Docker registry secret required to pull the image. The secret must already exist.",
		"path":            "Path defines the path to disk file in the container",
		"imagePullPolicy": "Image pull policy.\nOne of Always, Never, IfNotPresent.\nDefaults to Always if :latest tag is specified, or IfNotPresent otherwise.\nCannot be updated.\nMore info: https://kubernetes.io/docs/concepts/containers/images#updating-images\n+optional",
	}
}

func (ClockOffset) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "Exactly one of its members must be set.\n\n+k8s:openapi-gen=true",
		"utc":      "UTC sets the guest clock to UTC on each boot. If an offset is specified,\nguest changes to the clock will be kept during reboots and are not reset.",
		"timezone": "Timezone sets the guest clock to the specified timezone.\nZone name follows the TZ environment variable format (e.g. 'America/New_York').",
	}
}

func (ClockOffsetUTC) SwaggerDoc() map[string]string {
	return map[string]string{
		"":              "UTC sets the guest clock to UTC on each boot.\n\n+k8s:openapi-gen=true",
		"offsetSeconds": "OffsetSeconds specifies an offset in seconds, relative to UTC. If set,\nguest changes to the clock will be kept during reboots and not reset.",
	}
}

func (Clock) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "Represents the clock and timers of a vmi.\n\n+k8s:openapi-gen=true\n+kubebuilder:pruning:PreserveUnknownFields",
		"timer": "Timer specifies whih timers are attached to the vmi.\n+optional",
	}
}

func (Timer) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "Represents all available timers in a vmi.\n\n+k8s:openapi-gen=true",
		"hpet":   "HPET (High Precision Event Timer) - multiple timers with periodic interrupts.",
		"kvm":    "KVM \t(KVM clock) - lets guests read the host’s wall clock time (paravirtualized). For linux guests.",
		"pit":    "PIT (Programmable Interval Timer) - a timer with periodic interrupts.",
		"rtc":    "RTC (Real Time Clock) - a continuously running timer with periodic interrupts.",
		"hyperv": "Hyperv (Hypervclock) - lets guests read the host’s wall clock time (paravirtualized). For windows guests.",
	}
}

func (RTCTimer) SwaggerDoc() map[string]string {
	return map[string]string{
		"":           "+k8s:openapi-gen=true",
		"tickPolicy": "TickPolicy determines what happens when QEMU misses a deadline for injecting a tick to the guest.\nOne of \"delay\", \"catchup\".",
		"present":    "Enabled set to false makes sure that the machine type or a preset can't add the timer.\nDefaults to true.\n+optional",
		"track":      "Track the guest or the wall clock.",
	}
}

func (HPETTimer) SwaggerDoc() map[string]string {
	return map[string]string{
		"":           "+k8s:openapi-gen=true",
		"tickPolicy": "TickPolicy determines what happens when QEMU misses a deadline for injecting a tick to the guest.\nOne of \"delay\", \"catchup\", \"merge\", \"discard\".",
		"present":    "Enabled set to false makes sure that the machine type or a preset can't add the timer.\nDefaults to true.\n+optional",
	}
}

func (PITTimer) SwaggerDoc() map[string]string {
	return map[string]string{
		"":           "+k8s:openapi-gen=true",
		"tickPolicy": "TickPolicy determines what happens when QEMU misses a deadline for injecting a tick to the guest.\nOne of \"delay\", \"catchup\", \"discard\".",
		"present":    "Enabled set to false makes sure that the machine type or a preset can't add the timer.\nDefaults to true.\n+optional",
	}
}

func (KVMTimer) SwaggerDoc() map[string]string {
	return map[string]string{
		"":        "+k8s:openapi-gen=true",
		"present": "Enabled set to false makes sure that the machine type or a preset can't add the timer.\nDefaults to true.\n+optional",
	}
}

func (HypervTimer) SwaggerDoc() map[string]string {
	return map[string]string{
		"":        "+k8s:openapi-gen=true",
		"present": "Enabled set to false makes sure that the machine type or a preset can't add the timer.\nDefaults to true.\n+optional",
	}
}

func (Features) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "+k8s:openapi-gen=true",
		"acpi":   "ACPI enables/disables ACPI inside the guest.\nDefaults to enabled.\n+optional",
		"apic":   "Defaults to the machine type setting.\n+optional",
		"hyperv": "Defaults to the machine type setting.\n+optional",
		"smm":    "SMM enables/disables System Management Mode.\nTSEG not yet implemented.\n+optional",
		"kvm":    "Configure how KVM presence is exposed to the guest.\n+optional",
	}
}

func (FeatureState) SwaggerDoc() map[string]string {
	return map[string]string{
		"":        "Represents if a feature is enabled or disabled.\n\n+k8s:openapi-gen=true",
		"enabled": "Enabled determines if the feature should be enabled or disabled on the guest.\nDefaults to true.\n+optional",
	}
}

func (FeatureAPIC) SwaggerDoc() map[string]string {
	return map[string]string{
		"":               "+k8s:openapi-gen=true",
		"enabled":        "Enabled determines if the feature should be enabled or disabled on the guest.\nDefaults to true.\n+optional",
		"endOfInterrupt": "EndOfInterrupt enables the end of interrupt notification in the guest.\nDefaults to false.\n+optional",
	}
}

func (FeatureSpinlocks) SwaggerDoc() map[string]string {
	return map[string]string{
		"":          "+k8s:openapi-gen=true",
		"enabled":   "Enabled determines if the feature should be enabled or disabled on the guest.\nDefaults to true.\n+optional",
		"spinlocks": "Retries indicates the number of retries.\nMust be a value greater or equal 4096.\nDefaults to 4096.\n+optional",
	}
}

func (FeatureVendorID) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "+k8s:openapi-gen=true",
		"enabled":  "Enabled determines if the feature should be enabled or disabled on the guest.\nDefaults to true.\n+optional",
		"vendorid": "VendorID sets the hypervisor vendor id, visible to the vmi.\nString up to twelve characters.",
	}
}

func (FeatureHyperv) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                "Hyperv specific features.\n\n+k8s:openapi-gen=true",
		"relaxed":         "Relaxed instructs the guest OS to disable watchdog timeouts.\nDefaults to the machine type setting.\n+optional",
		"vapic":           "VAPIC improves the paravirtualized handling of interrupts.\nDefaults to the machine type setting.\n+optional",
		"spinlocks":       "Spinlocks allows to configure the spinlock retry attempts.\n+optional",
		"vpindex":         "VPIndex enables the Virtual Processor Index to help windows identifying virtual processors.\nDefaults to the machine type setting.\n+optional",
		"runtime":         "Runtime improves the time accounting to improve scheduling in the guest.\nDefaults to the machine type setting.\n+optional",
		"synic":           "SyNIC enables the Synthetic Interrupt Controller.\nDefaults to the machine type setting.\n+optional",
		"synictimer":      "SyNICTimer enables Synthetic Interrupt Controller Timers, reducing CPU load.\nDefaults to the machine type setting.\n+optional",
		"reset":           "Reset enables Hyperv reboot/reset for the vmi. Requires synic.\nDefaults to the machine type setting.\n+optional",
		"vendorid":        "VendorID allows setting the hypervisor vendor id.\nDefaults to the machine type setting.\n+optional",
		"frequencies":     "Frequencies improves the TSC clock source handling for Hyper-V on KVM.\nDefaults to the machine type setting.\n+optional",
		"reenlightenment": "Reenlightenment enables the notifications on TSC frequency changes.\nDefaults to the machine type setting.\n+optional",
		"tlbflush":        "TLBFlush improves performances in overcommited environments. Requires vpindex.\nDefaults to the machine type setting.\n+optional",
		"ipi":             "IPI improves performances in overcommited environments. Requires vpindex.\nDefaults to the machine type setting.\n+optional",
		"evmcs":           "EVMCS Speeds up L2 vmexits, but disables other virtualization features. Requires vapic.\nDefaults to the machine type setting.\n+optional",
	}
}

func (FeatureKVM) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "+k8s:openapi-gen=true",
		"hidden": "Hide the KVM hypervisor from standard MSR based discovery.\nDefaults to false",
	}
}

func (Watchdog) SwaggerDoc() map[string]string {
	return map[string]string{
		"":     "Named watchdog device.\n\n+k8s:openapi-gen=true",
		"name": "Name of the watchdog.",
	}
}

func (WatchdogDevice) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "Hardware watchdog device.\nExactly one of its members must be set.\n\n+k8s:openapi-gen=true",
		"i6300esb": "i6300esb watchdog device.\n+optional",
	}
}

func (I6300ESBWatchdog) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "i6300esb watchdog device.\n\n+k8s:openapi-gen=true",
		"action": "The action to take. Valid values are poweroff, reset, shutdown.\nDefaults to reset.",
	}
}

func (Interface) SwaggerDoc() map[string]string {
	return map[string]string{
		"":            "+k8s:openapi-gen=true",
		"name":        "Logical name of the interface as well as a reference to the associated networks.\nMust match the Name of a Network.",
		"model":       "Interface model.\nOne of: e1000, e1000e, ne2k_pci, pcnet, rtl8139, virtio.\nDefaults to virtio.",
		"ports":       "List of ports to be forwarded to the virtual machine.",
		"macAddress":  "Interface MAC address. For example: de:ad:00:00:be:af or DE-AD-00-00-BE-AF.",
		"bootOrder":   "BootOrder is an integer value > 0, used to determine ordering of boot devices.\nLower values take precedence.\nEach interface or disk that has a boot order must have a unique value.\nInterfaces without a boot order are not tried.\n+optional",
		"pciAddress":  "If specified, the virtual network interface will be placed on the guests pci address with the specifed PCI address. For example: 0000:81:01.10\n+optional",
		"dhcpOptions": "If specified the network interface will pass additional DHCP options to the VMI\n+optional",
		"tag":         "If specified, the virtual network interface address and its tag will be provided to the guest via config drive\n+optional",
	}
}

func (DHCPOptions) SwaggerDoc() map[string]string {
	return map[string]string{
		"":               "Extra DHCP options to use in the interface.\n\n+k8s:openapi-gen=true",
		"bootFileName":   "If specified will pass option 67 to interface's DHCP server\n+optional",
		"tftpServerName": "If specified will pass option 66 to interface's DHCP server\n+optional",
		"ntpServers":     "If specified will pass the configured NTP server to the VM via DHCP option 042.\n+optional",
		"privateOptions": "If specified will pass extra DHCP options for private use, range: 224-254\n+optional",
	}
}

func (DHCPPrivateOptions) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "DHCPExtraOptions defines Extra DHCP options for a VM.\n\n+k8s:openapi-gen=true",
		"option": "Option is an Integer value from 224-254\nRequired.",
		"value":  "Value is a String value for the Option provided\nRequired.",
	}
}

func (InterfaceBindingMethod) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "Represents the method which will be used to connect the interface to the guest.\nOnly one of its members may be specified.\n\n+k8s:openapi-gen=true",
	}
}

func (InterfaceBridge) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "+k8s:openapi-gen=true",
	}
}

func (InterfaceSlirp) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "+k8s:openapi-gen=true",
	}
}

func (InterfaceMasquerade) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "+k8s:openapi-gen=true",
	}
}

func (InterfaceSRIOV) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "+k8s:openapi-gen=true",
	}
}

func (InterfaceMacvtap) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "+k8s:openapi-gen=true",
	}
}

func (Port) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "Port repesents a port to expose from the virtual machine.\nDefault protocol TCP.\nThe port field is mandatory\n\n+k8s:openapi-gen=true",
		"name":     "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each\nnamed port in a pod must have a unique name. Name for the port that can be\nreferred to by services.\n+optional",
		"protocol": "Protocol for port. Must be UDP or TCP.\nDefaults to \"TCP\".\n+optional",
		"port":     "Number of port to expose for the virtual machine.\nThis must be a valid port number, 0 < x < 65536.",
	}
}

func (AccessCredentialSecretSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":           "+k8s:openapi-gen=true",
		"secretName": "SecretName represents the name of the secret in the VMI's namespace",
	}
}

func (ConfigDriveAccessCredentialPropagation) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "+k8s:openapi-gen=true",
	}
}

func (AuthorizedKeysFile) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "AuthorizedKeysFile represents a path within the guest\nthat ssh public keys should be propagated to\n\n+k8s:openapi-gen=true",
		"filePath": "FilePath represents the place on the guest that the authorized_keys\nfile should be writen to. This is expected to be a full path including\nboth the base directory and file name.",
	}
}

func (QemuGuestAgentAccessCredentialPropagation) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                    "+k8s:openapi-gen=true",
		"authorizedKeysFiles": "AuthorizedKeysFiles represents all the paths within the guest\nthat ssh public keys should be propagated to\n+optional\n+listType=set",
	}
}

func (SSHPublicKeyAccessCredentialSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "SSHPublicKeyAccessCredentialSource represents where to retrieve the ssh key\ncredentials\nOnly one of its members may be specified.\n\n+k8s:openapi-gen=true",
		"secret": "Secret means that the access credential is pulled from a kubernetes secret",
	}
}

func (SSHPublicKeyAccessCredentialPropagationMethod) SwaggerDoc() map[string]string {
	return map[string]string{
		"":               "SSHPublicKeyAccessCredentialPropagationMethod represents the method used to\ninject a ssh public key into the vm guest.\nOnly one of its members may be specified.\n\n+k8s:openapi-gen=true",
		"configDrive":    "ConfigDrivePropagation means that the ssh public keys are injected\ninto the VM using metadata using the configDrive cloud-init provider",
		"qemuGuestAgent": "QemuGuestAgentAccessCredentailPropagation means ssh public keys are\ndynamically injected into the vm at runtime via the qemu guest agent.\nThis feature requires the qemu guest agent to be running within the guest.",
	}
}

func (SSHPublicKeyAccessCredential) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                  "SSHPublicKeyAccessCredential represents a source and propagation method for\ninjecting ssh public keys into a vm guest\nOnly one of its members may be specified.\n\n+k8s:openapi-gen=true",
		"source":            "Source represents where the public keys are pulled from",
		"propagationMethod": "propagationMethod represents how the public key is injected into the vm guest.",
	}
}

func (AccessCredential) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "AccessCredential represents a credential source that can be used to\nauthorize remote access to the vm guest\nOnly one of its members may be specified.\n\n+k8s:openapi-gen=true",
	}
}

func (Network) SwaggerDoc() map[string]string {
	return map[string]string{
		"":     "Network represents a network type and a resource that should be connected to the vm.\n\n+k8s:openapi-gen=true",
		"name": "Network name.\nMust be a DNS_LABEL and unique within the vm.\nMore info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
	}
}

func (NetworkSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "Represents the source resource that will be connected to the vm.\nOnly one of its members may be specified.\n\n+k8s:openapi-gen=true",
	}
}

func (PodNetwork) SwaggerDoc() map[string]string {
	return map[string]string{
		"":              "Represents the stock pod network interface.\n\n+k8s:openapi-gen=true",
		"vmNetworkCIDR": "CIDR for vm network.\nDefault 10.0.2.0/24 if not specified.",
	}
}

func (Rng) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "Rng represents the random device passed from host\n\n+k8s:openapi-gen=true",
	}
}

func (MultusNetwork) SwaggerDoc() map[string]string {
	return map[string]string{
		"":            "Represents the multus cni network.\n\n+k8s:openapi-gen=true",
		"networkName": "References to a NetworkAttachmentDefinition CRD object. Format:\n<networkName>, <namespace>/<networkName>. If namespace is not\nspecified, VMI namespace is assumed.",
		"default":     "Select the default network and add it to the\nmultus-cni.io/default-network annotation.",
	}
}
