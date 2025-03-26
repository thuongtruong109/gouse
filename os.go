package gouse

import (
	"log"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

// const megabyte uint64 = 1024 * 1024
// const gigabyte uint64 = 1024 * megabyte

type ICPU struct {
	CPU        int32
	Speed      string
	Cores      int
	VendorID   string
	Family     string
	Model      string
	Stepping   int32
	PhysicalID string
	CoreID     string
	ModelName  string
	Mhz        float64
	CacheSize  int32
	Flags      []string
	Microcode  string
}

func CPU() (*ICPU, error) {
	cpuStat, err := cpu.Info()
	if err != nil {
		return &ICPU{}, err
	}

	return &ICPU{
		CPU:        cpuStat[0].CPU,
		VendorID:   cpuStat[0].VendorID,
		Model:      cpuStat[0].Model,
		Stepping:   cpuStat[0].Stepping,
		PhysicalID: cpuStat[0].PhysicalID,
		CoreID:     cpuStat[0].CoreID,
		ModelName:  cpuStat[0].ModelName,
		Family:     cpuStat[0].Family,
		Speed:      strconv.FormatFloat(cpuStat[0].Mhz, 'f', 2, 64),
		Cores:      runtime.NumCPU(),
		Mhz:        cpuStat[0].Mhz,
		CacheSize:  cpuStat[0].CacheSize,
		Flags:      cpuStat[0].Flags,
		Microcode:  cpuStat[0].Microcode,
	}, nil
}

type IDisk struct {
	Path              string
	Fstype            string
	Total             uint64
	Free              uint64
	Used              uint64
	UsedPercent       float64
	InodesTotal       uint64
	InodesUsed        uint64
	InodesFree        uint64
	InodesUsedPercent float64
}

func Disk() (*IDisk, error) {
	diskStat, err := disk.Usage("/")
	if err != nil {
		return &IDisk{}, err
	}

	return &IDisk{
		Total:             diskStat.Total,
		Free:              diskStat.Free,
		Used:              diskStat.Used,
		UsedPercent:       diskStat.UsedPercent, // for string type: strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64),
		InodesTotal:       diskStat.InodesTotal,
		InodesUsed:        diskStat.InodesUsed,
		InodesFree:        diskStat.InodesFree,
		InodesUsedPercent: diskStat.InodesUsedPercent,
	}, nil
}

type IIo struct {
	ReadCount        uint64
	MergedReadCount  uint64
	WriteCount       uint64
	MergedWriteCount uint64
	ReadBytes        uint64
	WriteBytes       uint64
	ReadTime         uint64
	WriteTime        uint64
	IopsInProgress   uint64
	IoTime           uint64
	WeightedIO       uint64
	Name             string
	SerialNumber     string
	Label            string
}

func Io() ([]*IIo, error) {
	ioStat, err := disk.IOCounters()
	if err != nil {
		return nil, err
	}

	var io []*IIo
	for _, v := range ioStat {
		io = append(io, &IIo{
			ReadCount:        v.ReadCount,
			MergedReadCount:  v.MergedReadCount,
			WriteCount:       v.WriteCount,
			MergedWriteCount: v.MergedWriteCount,
			ReadBytes:        v.ReadBytes,
			WriteBytes:       v.WriteBytes,
			ReadTime:         v.ReadTime,
			WriteTime:        v.WriteTime,
			IopsInProgress:   v.IopsInProgress,
			IoTime:           v.IoTime,
			WeightedIO:       v.WeightedIO,
			Name:             v.Name,
			SerialNumber:     v.SerialNumber,
			Label:            v.Label,
		})
	}

	return io, nil
}

type IPartition struct {
	Device     string
	Mountpoint string
	Fstype     string
	Opts       []string
}

func Partition() ([]*IPartition, error) {
	partitionStat, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}

	var partitions []*IPartition
	for _, v := range partitionStat {
		partitions = append(partitions, &IPartition{
			Device:     v.Device,
			Mountpoint: v.Mountpoint,
			Fstype:     v.Fstype,
			Opts:       v.Opts,
		})
	}

	return partitions, nil
}

type IHost struct {
	Name                 string
	Platform             string
	Arch                 string
	Hostname             string
	NumsProcs            uint64
	Uptime               uint64
	BootTime             uint64
	Procs                uint64
	OS                   string
	PlatformFamily       string
	PlatformVersion      string
	KernelVersion        string
	KernelArch           string
	VirtualizationSystem string
	VirtualizationRole   string
	HostID               string
}

func Host() (*IHost, error) {
	runTimeOS := runtime.GOOS

	hostStat, err := host.Info()
	if err != nil {
		return &IHost{}, err
	}

	return &IHost{
		Name:                 runTimeOS,
		Platform:             hostStat.Platform,
		Arch:                 hostStat.KernelArch,
		Hostname:             hostStat.Hostname,
		NumsProcs:            hostStat.Procs,
		Uptime:               hostStat.Uptime,
		BootTime:             hostStat.BootTime,
		Procs:                hostStat.Procs,
		OS:                   hostStat.OS,
		PlatformFamily:       hostStat.PlatformFamily,
		PlatformVersion:      hostStat.PlatformVersion,
		KernelVersion:        hostStat.KernelVersion,
		KernelArch:           hostStat.KernelArch,
		VirtualizationSystem: hostStat.VirtualizationSystem,
		VirtualizationRole:   hostStat.VirtualizationRole,
		HostID:               hostStat.HostID,
	}, nil
}

type IUser struct {
	User     string
	Terminal string
	Host     string
	Started  int
}

func User() ([]*IUser, error) {
	userStat, err := host.Users()
	if err != nil {
		return nil, err
	}

	var users []*IUser
	for _, v := range userStat {
		users = append(users, &IUser{
			User:     v.User,
			Terminal: v.Terminal,
			Host:     v.Host,
			Started:  v.Started,
		})
	}

	return users, nil
}

type IMemory struct {
	VirtualTotal          uint64
	VirtualAvailable      uint64
	VirtualUsed           uint64
	VirtualUsedPercent    float64
	VirtualFree           uint64
	VirtualActive         uint64
	VirtualInactive       uint64
	VirtualWired          uint64
	VirtualLaundry        uint64
	VirtualBuffers        uint64
	VirtualCached         uint64
	VirtualWriteBack      uint64
	VirtualDirty          uint64
	VirtualWriteBackTmp   uint64
	VirtualShared         uint64
	VirtualSlab           uint64
	VirtualSreclaimable   uint64
	VirtualSunreclaim     uint64
	VirtualPageTables     uint64
	VirtualSwapCached     uint64
	VirtualCommitLimit    uint64
	VirtualCommittedAS    uint64
	VirtualHighTotal      uint64
	VirtualHighFree       uint64
	VirtualLowTotal       uint64
	VirtualLowFree        uint64
	VirtualMapped         uint64
	VirtualMallocTotal    uint64
	VirtualMallocUsed     uint64
	VirtualMallocChunk    uint64
	VirtualHugePagesTotal uint64
	VirtualHugePagesFree  uint64
	VirtualHugePagesRsvd  uint64
	VirtualHugePagesSurp  uint64
	VirtualHugePageSize   uint64
	VirtualAnonHugePages  uint64

	SwapTotal       uint64
	SwapUsed        uint64
	SwapFree        uint64
	SwapUsedPercent float64
	SwapSin         uint64
	SwapSout        uint64
	SwapPgIn        uint64
	SwapPgOut       uint64
	SwapPgFault     uint64
	SwapPgMajFault  uint64
}

func Memory() (*IMemory, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return &IMemory{}, err
	}

	swapStat, err := mem.SwapMemory()
	if err != nil {
		return &IMemory{}, err
	}

	return &IMemory{
		VirtualTotal:          vmStat.Total,
		VirtualAvailable:      vmStat.Available,
		VirtualUsed:           vmStat.Used,
		VirtualUsedPercent:    vmStat.UsedPercent, // for string type: strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64),
		VirtualFree:           vmStat.Free,
		VirtualActive:         vmStat.Active,
		VirtualInactive:       vmStat.Inactive,
		VirtualWired:          vmStat.Wired,
		VirtualLaundry:        vmStat.Laundry,
		VirtualBuffers:        vmStat.Buffers,
		VirtualCached:         vmStat.Cached,
		VirtualWriteBack:      vmStat.WriteBack,
		VirtualDirty:          vmStat.Dirty,
		VirtualWriteBackTmp:   vmStat.WriteBackTmp,
		VirtualShared:         vmStat.Shared,
		VirtualSlab:           vmStat.Slab,
		VirtualSreclaimable:   vmStat.Sreclaimable,
		VirtualSunreclaim:     vmStat.Sunreclaim,
		VirtualPageTables:     vmStat.PageTables,
		VirtualSwapCached:     vmStat.SwapCached,
		VirtualCommitLimit:    vmStat.CommitLimit,
		VirtualCommittedAS:    vmStat.CommittedAS,
		VirtualHighTotal:      vmStat.HighTotal,
		VirtualHighFree:       vmStat.HighFree,
		VirtualLowTotal:       vmStat.LowTotal,
		VirtualLowFree:        vmStat.LowFree,
		VirtualMapped:         vmStat.Mapped,
		VirtualMallocTotal:    vmStat.VmallocTotal,
		VirtualMallocUsed:     vmStat.VmallocUsed,
		VirtualMallocChunk:    vmStat.VmallocChunk,
		VirtualHugePagesTotal: vmStat.HugePagesTotal,
		VirtualHugePagesFree:  vmStat.HugePagesFree,
		VirtualHugePagesRsvd:  vmStat.HugePagesRsvd,
		VirtualHugePagesSurp:  vmStat.HugePagesSurp,
		VirtualHugePageSize:   vmStat.HugePageSize,
		VirtualAnonHugePages:  vmStat.AnonHugePages,

		SwapTotal:       swapStat.Total,
		SwapUsed:        swapStat.Used,
		SwapFree:        swapStat.Free,
		SwapUsedPercent: swapStat.UsedPercent,
		SwapSin:         swapStat.Sin,
		SwapSout:        swapStat.Sout,
		SwapPgIn:        swapStat.PgIn,
		SwapPgOut:       swapStat.PgOut,
		SwapPgFault:     swapStat.PgFault,
		SwapPgMajFault:  swapStat.PgMajFault,
	}, nil
}

func Profile(cpuprofile, memprofile string) {
	closeFile := func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Printf("error closing profile file: %v", err)
		}
	}

	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer closeFile(f)

		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}
	// ... rest of the program ...
	if memprofile != "" {
		f, err := os.Create(memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer closeFile(f)

		runtime.GC()

		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}