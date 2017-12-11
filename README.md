ynniogwen-proxy
=================

Scrapes values from the Ynni Ogwen turbine and serves on HTTP.

Building
--------

Dependecies are managed with [dep](https://github.com/golang/dep)
```bash
dep ensure
cd ynniogwen-proxy
go build
```

Running
-------

    ./ynniogwen-proxy -url http://1.2.3.4/

HTTP Routes
------
[http://localhost:9000/](http://localhost:9000/) JSON (cached for 1 minute) e.g.
```json
{
    "G59 mains OK": 1,
    "GV1 position": 42,
    "GV2 position": 64,
    "Gearbox Temperature": 56,
    "P1 Mains Voltage": 249,
    "Power Output": 62,
    "River Level": 143,
    "Softstarter Closed": 1,
    "Speed": 399,
    "Sump Level": -106,
    "Theoretical Power": 68,
    "Time (hours)": 0,
    "Time (mins)": 9,
    "kWh": 719,
    "mWh": 329
}
```

[http://localhost:9000/metrics](http://localhost:9000/metrics) Endpoint for [Prometheus](https://prometheus.io/) e.g.
```
# HELP go_gc_duration_seconds A summary of the GC invocation durations.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 0
go_gc_duration_seconds{quantile="0.25"} 0
go_gc_duration_seconds{quantile="0.5"} 0
go_gc_duration_seconds{quantile="0.75"} 0
go_gc_duration_seconds{quantile="1"} 0
go_gc_duration_seconds_sum 0
go_gc_duration_seconds_count 0
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 14
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 615800
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 615800
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 2758
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 311
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 137216
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 615800
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 73728
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 1.564672e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 6387
# HELP go_memstats_heap_released_bytes_total Total number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes_total counter
go_memstats_heap_released_bytes_total 0
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 1.6384e+06
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 0
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 12
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 6698
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 13888
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 16384
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 25992
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 32768
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.473924e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 798010
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 458752
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 458752
# HELP go_memstats_sys_bytes Number of bytes obtained by system. Sum of all system allocations.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 3.084288e+06
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 0
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 1024
# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge
process_open_fds 7
# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 6.959104e+06
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.51295108871e+09
# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 3.42429696e+08
# HELP ynni_ogwen_turbine_g59_mains_ok Mains OK (boolean).
# TYPE ynni_ogwen_turbine_g59_mains_ok gauge
ynni_ogwen_turbine_g59_mains_ok 1
# HELP ynni_ogwen_turbine_gearbox_temperature Gearbox temperature (Â°C)
# TYPE ynni_ogwen_turbine_gearbox_temperature gauge
ynni_ogwen_turbine_gearbox_temperature 56
# HELP ynni_ogwen_turbine_gv1_position GV1 position.
# TYPE ynni_ogwen_turbine_gv1_position gauge
ynni_ogwen_turbine_gv1_position 42
# HELP ynni_ogwen_turbine_gv2_position GV2 position.
# TYPE ynni_ogwen_turbine_gv2_position gauge
ynni_ogwen_turbine_gv2_position 64
# HELP ynni_ogwen_turbine_kWh Kilowatt hours.
# TYPE ynni_ogwen_turbine_kWh gauge
ynni_ogwen_turbine_kWh 719
# HELP ynni_ogwen_turbine_mWh Megawatt hours.
# TYPE ynni_ogwen_turbine_mWh gauge
ynni_ogwen_turbine_mWh 329
# HELP ynni_ogwen_turbine_p1_mains_voltage Mains Voltage.
# TYPE ynni_ogwen_turbine_p1_mains_voltage gauge
ynni_ogwen_turbine_p1_mains_voltage 249
# HELP ynni_ogwen_turbine_power_output Power output of the turbine (KWh).
# TYPE ynni_ogwen_turbine_power_output gauge
ynni_ogwen_turbine_power_output 62
# HELP ynni_ogwen_turbine_river_level River level (mm?)
# TYPE ynni_ogwen_turbine_river_level gauge
ynni_ogwen_turbine_river_level 143
# HELP ynni_ogwen_turbine_softstarted_closed Soft starter closed (boolean).
# TYPE ynni_ogwen_turbine_softstarted_closed gauge
ynni_ogwen_turbine_softstarted_closed 1
# HELP ynni_ogwen_turbine_speed Speed (RPM?)
# TYPE ynni_ogwen_turbine_speed gauge
ynni_ogwen_turbine_speed 399
# HELP ynni_ogwen_turbine_sump_level Sump level (unit unknown)
# TYPE ynni_ogwen_turbine_sump_level gauge
ynni_ogwen_turbine_sump_level -106
# HELP ynni_ogwen_turbine_theoretical_power Theoretical power output (KWh).
# TYPE ynni_ogwen_turbine_theoretical_power gauge
ynni_ogwen_turbine_theoretical_power 68
```
