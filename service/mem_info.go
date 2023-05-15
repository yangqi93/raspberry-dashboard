package service

import (
	"math"
	"os"
	"regexp"
	"strconv"
)

func MemInfo() (memory *Mem, err error) {
	mem, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return nil, err
	}
	str := string(mem)
	//正则匹配
	reg := regexp.MustCompile(`MemTotal\:.*?([\d\.]+).+?\n`).FindStringSubmatch(str)
	var total, free, buffers, cached, cached_percent, used, percent, real_used, real_free, real_percent, swap_total, swap_free, swap_used, swap_percent float64
	if len(reg) >= 2 {
		i, _ := strconv.Atoi(reg[1])
		total = math.Round(float64(i / 1024))
	}
	reg = regexp.MustCompile(`MemFree\:.*?([\d\.]+).+?\n`).FindStringSubmatch(str)
	if len(reg) >= 2 {
		i, _ := strconv.Atoi(reg[1])
		free = math.Round(float64(i / 1024))
	}
	reg = regexp.MustCompile(`Buffers\:.*?([\d\.]+).+?\n`).FindStringSubmatch(str)
	if len(reg) >= 2 {
		i, _ := strconv.Atoi(reg[1])
		buffers = math.Round(float64(i / 1024))
	}
	reg = regexp.MustCompile(`Cached\:.*?([\d\.]+).+?\n`).FindStringSubmatch(str)
	if len(reg) >= 2 {
		i, _ := strconv.Atoi(reg[1])
		cached = math.Round(float64(i / 1024))
	}
	reg = regexp.MustCompile(`SwapTotal\:.*?([\d\.]+).+?\n`).FindStringSubmatch(str)
	if len(reg) >= 2 {
		i, _ := strconv.Atoi(reg[1])
		swap_total = math.Round(float64(i / 1024))
	}
	reg = regexp.MustCompile(`SwapFree\:.*?([\d\.]+).+?\n`).FindStringSubmatch(str)
	if len(reg) >= 2 {
		i, _ := strconv.Atoi(reg[1])
		swap_free = math.Round(float64(i / 1024))
	}
	// cached percent
	if cached != 0 && total != 0 {
		cached_percent = math.Round(cached / total * 100)
	}
	// used
	if total != 0 && free != 0 {
		used = total - free
	}
	// percent
	if total != 0 && used != 0 {
		percent = math.Round(used / total * 100)
	}
	// real_used
	if total != 0 && free != 0 && buffers != 0 && cached != 0 {
		real_used = total - free - buffers - cached
	}
	// real_free
	if total != 0 && real_used != 0 {
		real_free = math.Round(total - real_used)
	}
	// real_percent
	if total != 0 && real_used != 0 {
		real_percent = math.Round(real_used / total * 100)
	}
	// swap_used
	if swap_total != 0 && swap_free != 0 {
		swap_used = math.Round(swap_total - swap_free)
	}
	// swap_percent
	if swap_total != 0 && swap_used != 0 {
		swap_percent = math.Round(swap_used / swap_total * 100)
	}

	memory.Total = total
	memory.Free = free
	memory.Buffers = buffers
	memory.Cached = cached
	memory.CachedPercent = cached_percent
	memory.Used = used
	memory.Percent = percent
	memory.Real.Used = real_used
	memory.Real.Free = real_free
	memory.Real.Percent = real_percent
	memory.Swap.Total = int(swap_total)
	memory.Swap.Free = int(swap_free)
	memory.Swap.Used = int(swap_used)
	memory.Swap.Percent = int(swap_percent)

	return memory, nil
}
