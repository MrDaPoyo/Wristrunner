# PX30 Release Note

## px30_ddr_333MHz_wdqs_v2.12.bin

| Date       | file                           | Build commit | Severity  |
| ---------- | ------------------------------ | ------------ | --------- |
| 2025-11-20 | px30_ddr_333MHz_wdqs_v2.12.bin | 2a37520      | important |

### Warn

1. When there is no DQSn pull-up resistor on the PCB, only the regular DDRBIN in the bin/rk33 directory can be used. When there is a DQSn pull-up resistor on the PCB, only this version of DDRBIN can be used. Otherwise, there may be reliability issues.

### New

1. Add WDQS support. And a 560 ohm weak pull-up resistor is required on DQSn on the PCB.

------

