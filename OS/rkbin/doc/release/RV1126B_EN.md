# rv1126b Release Note

## rv1126b_usbplug_v1.03.bin

| Date       | File                       | Build commit | Severity  |
| ---------- | :------------------------- | ------------ | --------- |
| 2026-06-17 | rv1126b_usbplug_v1.03.bin  | df65f0e1b4   | important |

### New

1. Add new spiflash support.

------

## rv1126b_bl32_v1.06.bin

| Date       | File                   | Build Commit | Severity  |
| ---------- | ---------------------- | ------------ | --------- |
| 2026-04-29 | rv1126b_bl32_v1.06.bin | 5f256dc1a    | important |

### Fixed

| Index | Severity  | Update                                   | Issue description                                            | Issue source |
| ----- | --------- | ---------------------------------------- | ------------------------------------------------------------ | ------------ |
| 1     | important | Fix the low-probability data abort issue | The reboot of the device during secure storage may trigger it with a low probability | -            |

------

## rv1126b_bl32_v1.05.bin

| Date       | File                   | Build Commit | Severity |
| ---------- | ---------------------- | ------------ | -------- |
| 2026-03-10 | rv1126b_bl32_v1.05.bin | 14421695d    | important |

### New

1. Support CTR mode.
2. OAEP decode compatible with MGF1 = SHA1.
3. Support DICE data read/write.

### Fixed

| Index | Severity | Update                 | Issue description                 | Issue source |
| ----- | -------- | ---------------------------------- | ---------------------------------------- | -------- |
| 1     | important | Enable CE instruction set by default | Algorithm runs slow without CE support | 614919 |

------

## rv1126b_spl{, _ipc}_v1.06.bin

| Date       | File                  | Build commit | Severity  |
| ---------- | :----------------------- | ----------- | -------- |
| 2026-03-02 | rv1126b_spl{, _ipc}_v1.06.bin | a27a89c5fa1   | important     |

### Fixed

| Index | Severity  | Update                                  | Issue description                   | Issue source |
| ----- | --------- | --------------------------------------- | ----------------------------------- | ------------ |
| 1     | important  | Avoid premature decompression termination by decom. | SPL hw decompression of uboot failed. | -            |

------

## rv1126b{p}_ddr_{1332, 1056}MHz_v1.10.bin

| Date       | File                                     | Build commit | Severity |
| ---------- | :--------------------------------------- | ------------ | -------- |
| 2026-01-06 | rv1126b{p}_ddr_{1332, 1056}MHz_v1.10.bin | 53750e2af0   | moderate |

### New

1. Support using the tool to change the RV1126BP frequency to a maximum of 1332MHz. Note: Please confirm that the hardware design can achieve the target frequency before making the change.

2. Compatible with some UniIC and ISSI LPDDR4(X).

3. Add rv1126b_ddr_Template_DDR3P216DD61346_32R5x25R0_2112Mbps_H1R6_V10_20251212_1056MHz_{tb_, eyescan_}v1.10.bin, only for supporting the template: RV1126B_Template_DDR3P216DD61346_32R5x25R0_2112Mbps_H1R6_V10_20251212.

------

## rv1126b_bl31_v1.13.elf

| Date       | File                  | Build commit | Severity  |
| ---------- | :-------------------- | ------------ | --------- |
| 2025-12-10 | rv1126b_bl31_v1.13.elf | be86983f3 | important |

### New

1. Improve the frequency accuracy of RC OSC when system suspend.

------

## rv1126b{p}_ddr_{1332, 1056}MHz_v1.09.bin

| Date       | File                                     | Build commit | Severity |
| ---------- | :--------------------------------------- | ------------ | -------- |
| 2025-11-27 | rv1126b{p}_ddr_{1332, 1056}MHz_v1.09.bin | 6d04c3b809   | moderate |

### Fixed

| Index | Severity | Update                                                       | Issue description | Issue source |
| ----- | -------- | ------------------------------------------------------------ | ----------------- | ------------ |
| 1     | moderate | Resolve the issue of VOP screen flickering in some scenarios under DDR total bit width of 16bit. | -                 | -            |

------

## rv1126b{p}_ddr_{1332, 1056}MHz_v1.08.bin

| Date       | File                                     | Build commit | Severity  |
| ---------- | :--------------------------------------- | ------------ | --------- |
| 2025-11-20 | rv1126b{p}_ddr_{1332, 1056}MHz_v1.08.bin | 1cbce6fa94   | moderate |

### New

1. Support RV1126BM.

------

## rv1126b_bl31_v1.12.elf

| Date       | File                  | Build commit | Severity  |
| ---------- | :-------------------- | ------------ | --------- |
| 2025-11-04 | rv1126b_bl31_v1.12.elf | ae491a342 | important |

### New

1. Support 16bit LPDDR4/LPDDR4X for system suspend.
2. Optimize system resume time.

------

## rv1126b{p}_ddr_{1332, 1056}MHz_v1.07.bin

| Date       | File                                     | Build commit | Severity  |
| ---------- | :--------------------------------------- | ------------ | --------- |
| 2025-11-20 | rv1126b{p}_ddr_{1332, 1056}MHz_v1.07.bin | 776cd76d6e   | moderate |

### New

1. Support RV1126BP LPDDR4X.

------

## rv1126b_spl{, _ipc}_v1.05.bin

| Date       | File                   | Build commit | Severity  |
| ---------- | :--------------------- | ------------ | --------- |
| 2025-09-26 | rv1126b_spl{, _ipc}_v1.05.bin | 2f340fbbc13    | important |

### New

1. Adjust tsadc offset and bias according to otp.

------

## rv1126b_bl31_v1.11.elf

| Date       | File                  | Build commit | Severity  |
| ---------- | :-------------------- | ------------ | --------- |
| 2025-09-19 | rv1126b_bl31_v1.11.elf | 405ded103 | important |

### New

1. Support ultra-mode for system sleep.
2. Support use non-ECC mode to read OTP.

------

## rv1126b_spl{, _ipc}_v1.04.bin

| Date       | File                   | Build commit | Severity  |
| ---------- | :--------------------- | ------------ | --------- |
| 2025-09-11 | rv1126b_spl{, _ipc}_v1.04.bin | 23541e77b13    | important |

| Index | Severity  | Update             | Issue description                                            | Issue source |
| ----- | --------- | ------------------ | ------------------------------------------------------------ | ------------ |
| 1     | important | Fix default bias | Due to the width, some temperatures may overflow. | -            |

------

## rv1126b{p}_ddr_{1332, 1056}MHz_v1.06.bin

| Date       | File                                     | Build commit | Severity  |
| ---------- | :--------------------------------------- | ------------ | --------- |
| 2025-09-11 | rv1126b{p}_ddr_{1332, 1056}MHz_v1.06.bin | 1604bf9935   | important |

### New

1. Support RV1126BJ.
2. Support pstore.

### Fixed

| Index | Severity  | Update             | Issue description                                            | Issue source |
| ----- | --------- | ------------------ | ------------------------------------------------------------ | ------------ |
| 1     | important | Fix refresh_margin | Some LPDDR4(X) may be unstable at specific frequency under high temperatures | 572300       |

------

## rv1126b_bl31_v1.10.elf

| Date       | File                  | Build commit | Severity  |
| ---------- | :-------------------- | ------------ | --------- |
| 2025-09-05 | rv1126b_bl31_v1.10.elf | 33e45088f | important |

### New

1. Save and restore more registers when system suspend.

------

## rv1126b_bl31_v1.09.elf

| Date       | File                  | Build commit | Severity  |
| ---------- | :-------------------- | ------------ | --------- |
| 2025-08-25 | rv1126b_bl31_v1.09.elf | 5b828b3a5 | important |

### New

1. Support to use external 32k clock during system suspend.

------

## rv1126b{p}_ddr_{1332, 1056}MHz_v1.05.bin

| Date       | File                                     | Build commit | Severity |
| ---------- | :--------------------------------------- | ------------ | -------- |
| 2025-08-18 | rv1126b{p}_ddr_{1332, 1056}MHz_v1.05.bin | c1d02225c7   | moderate |

### Fixed

| Index | Severity | Update                                  | Issue description             | Issue source |
| ----- | -------- | --------------------------------------- | ----------------------------- | ------------ |
| 1     | moderate | Optimize CA training PLL configuration  | Enhance CA training stability |              |
| 2     | moderate | Optimize frequency information printing | -                             |              |

------

## rv1126b_bl32_v1.04.bin

| Date       | File                   | Build commit | Severity  |
| ---------- | :--------------------- | ------------ | --------- |
| 2025-08-15 | rv1126b_bl32_v1.04.bin | a89319de7    | important |

### New

1. OTP supports Non-Protected OEM Zone functionality.
2. Support CMAC KDF derived algorithms.

------

## rv1126b{p}_ddr_{1332, 1056}MHz_v1.04.bin

| Date       | File                                     | Build commit | Severity  |
| ---------- | :--------------------------------------- | ------------ | --------- |
| 2025-07-22 | rv1126b{p}_ddr_{1332, 1056}MHz_v1.04.bin | ff328eecdd   | important |

### Fixed

| Index | Severity  | Update                                    | Issue description                              | Issue source |
| ----- | --------- | ----------------------------------------- | ---------------------------------------------- | ------------ |
| 1     | important | Fix​​ per-channel 3Gb/4Gb LPDDR4(X) tRFCpb  | Per-channel 3Gb/4Gb LPDDR4(X) sleep failures   |              |
| 2     | important | Modify​​ CS training mode                   | Initialization issues on some LPDDR4(X) boards |              |
| 3     | important | Optimize​​ DDR3/DDR4 SI settings            |                                                |              |
| 4     | important | Enable​​ RV1126B LPDDR4 DBI to enhance SIPI |                                                |              |
| 5     | moderate  | Optimize​​ LPDDR4(X) vref settings          |                                                |              |
| 6     | moderate  | Add​​ PHY pll_ls configuration              |                                                |              |

------

## rv1126b_bl31_v1.08.elf

| Date       | File                  | Build commit | Severity  |
| ---------- | :-------------------- | ------------ | --------- |
| 2025-07-17 | rv1126b_bl31_v1.08.elf | 2905dde0d | important |

### New

1. Always route fiq from s_el1  to el3.
2. Improve accuracy of system timer after system suspend/resume for AOA.

------

## rv1126b_bl31_v1.07.elf

| Date       | File                  | Build commit | Severity  |
| ---------- | :-------------------- | ------------ | --------- |
| 2025-06-13 | rv1126b_bl31_v1.07.elf | 46bfa4121 | important |

### New

1. Save/restore PVTPLL when system suspend/resume.
2. Improve stability of system reboot.
3. Support to boot HPMCU.
4. Improve stability of system temperature control.

------

## rv1126b{p}_ddr_{1332, 1056}MHz_v1.03.bin

| Date       | File                                     | Build commit | Severity |
| ---------- | :--------------------------------------- | ------------ | -------- |
| 2025-06-13 | rv1126b{p}_ddr_{1332, 1056}MHz_v1.03.bin | c4d8224017   | critical |

### Warn

1. The bl31 should be updated to v1.06 or above.

### Fixed

| Index | Severity  | Update                    | Issue description                                            | Issue source |
| ----- | --------- | ------------------------- | ------------------------------------------------------------ | ------------ |
| 1     | important | Fix refresh_margin        | LPDDR4(X) and LPDDR3 may be unstable at elevated temperatures and low frequencies | -            |
| 2     | critical  | Fix DDR4 stability issues |                                                              | -            |

------

## rv1126b_bl31_v1.06.elf

| Date       | File                  | Build commit | Severity  |
| ---------- | :-------------------- | ------------ | --------- |
| 2025-06-06 | rv1126b_bl31_v1.06.elf | 1b7f2f7db | important |

### New

1. Improve the speed of system suspend/resume.
2. Enhance the stability of system suspend/resume for DDR4 boards.

------

## rv1126b_hpmcu_tb_v1.00.bin

| Date       | File                  | Build commit | Severity |
| ---------- | :-------------------- | ----------- | -------- |
| 2025-05-27 | rv1126b_hpmcu_tb_v1.00.bin | rtt:3377af5ff6#hal:5c682f97 | important |

### New

1. Initial version.

------

## rv1126b_bl31_v1.05.elf

| Date       | File                  | Build commit | Severity  |
| ---------- | :-------------------- | ------------ | --------- |
| 2025-05-16 | rv1126b_bl31_v1.05.elf | ce0a773ac | important |

### New

1. Support fiq-debugger.
2. Support lp_pr sleep mode.

------

## rv1126b_bl32_v1.03.bin

| Date       | File                   | Build commit | Severity  |
| ---------- | :--------------------- | ------------ | --------- |
| 2025-05-12 | rv1126b_bl32_v1.03.bin | 6fa7d182f    | important |

### Fixed

| Index | Severity  | Update                                                   | Issue description | Issue source |
| ----- | --------- | -------------------------------------------------------- | ----------------- | ------------ |
| 1     | important | fix crypto result error when using dynamic shared memory | -                 | -            |

------

## rv1126b_usbplug_v1.03.bin

| Date       | File                   | Build commit | Severity  |
| ---------- | :--------------------- | ------------ | --------- |
| 2025-05-13 | rv1126b_usbplug_v1.03.bin | 37d8770ee4   | important |

### New

1. Fix reboot fail.

------

## rv1126b_usbplug_v1.02.bin

| Date       | File                   | Build commit | Severity  |
| ---------- | :--------------------- | ------------ | --------- |
| 2025-05-09 | rv1126b_usbplug_v1.02.bin | ac302d626c    | important |

### New

1. Add uart upgrade support.

------

## rv1126b_spl{, _ipc}_v1.03.bin

| Date       | File                   | Build commit | Severity  |
| ---------- | :--------------------- | ------------ | --------- |
| 2025-05-08 | rv1126b_spl{, _ipc}_v1.03.bin | bc03ef0985f    | important |

### New

1. Fix bias error.

------

## rv1126b_bl31_v1.04.elf

| Date       | File                  | Build commit | Severity  |
| ---------- | :-------------------- | ------------ | --------- |
| 2025-05-06 | rv1126b_bl31_v1.04.elf | 886a0421f | important |

### New

1. Support LP_AOA sleep mode.
2. Support non-secure configuration remap_dsmc, remap_perixip, remap_pmuxip.
3. Enable tsadc_shut_m0 function.

------

## rv1126b_spl{, _ipc}_v1.02.bin

| Date       | File                   | Build commit | Severity  |
| ---------- | :--------------------- | ------------ | --------- |
| 2025-05-06 | rv1126b_spl{, _ipc}_v1.02.bin | e4069c1b773    | important |

### New

1. Adjust tsadc bias current.

------

## rv1126b{p}_ddr_{1332, 1056}MHz_v1.02.bin

| Date       | File                                     | Build commit | Severity |
| ---------- | :--------------------------------------- | ------------ | -------- |
| 2025-04-25 | rv1126b{p}_ddr_{1332, 1056}MHz_v1.02.bin | 701e251084   | critical |

### New

1. Added support for arch timer 1GHz.
2. Added thunder-boot DDR bin.
3. Added eyescan DDR bin.

### Fixed

| Index | Severity  | Update                                                       | Issue description             | Issue source |
| ----- | --------- | ------------------------------------------------------------ | ----------------------------- | ------------ |
| 1     | important | Optimized SI configuration                                   | -                             | -            |
| 2     | critical  | Fix clk skew calculation error in certain scenarios          | Instability in some DDR3/DDR4 | -            |
| 3     | critical  | Fix LPDDR4(X) stability issues                               | -                             | -            |
| 4     | critical  | Resolve initialization failures for specific memory capacity combinations | -                             | -            |

------

## rv1126b_bl32_v1.02.bin

| Date       | File                   | Build commit | Severity  |
| ---------- | :--------------------- | ------------ | --------- |
| 2025-04-23 | rv1126b_bl32_v1.02.bin | 106240c23    | important |

### Fixed

| Index | Severity  | Update                                  | Issue description | Issue source |
| ----- | --------- | --------------------------------------- | ----------------- | ------------ |
| 1     | important | correct whether reading ns otp with ecc | -                 | -            |

------

## rv1126b_bl31_v1.03.elf

| Date       | File                  | Build commit | Severity  |
| ---------- | :-------------------- | ------------ | --------- |
| 2025-04-16 | rv1126b_bl31_v1.03.elf | d93a2fceb | important |

### New

1. Improve stability of system suspend .

------

## rv1126b_bl31_v1.02.elf

| Date       | File                   | Build commit | Severity  |
| ---------- | :--------------------- | ------------ | --------- |
| 2025-04-15 | rv1126b_bl31_v1.02.elf | 32ea6e285    | important |

### New

1. BL31 runs at 0x40000000.
2. Save/restore data of system sram when system suspend/resume.

------

## rv1126b_spl_v1.01.bin

| Date       | File                   | Build commit | Severity  |
| ---------- | :--------------------- | ------------ | --------- |
| 2025-04-14 | rv1126b_spl_v1.01.bin | 8532419ba16    | important |

### New

1. Add arch timer 1Ghz support.

------

## rv1126b_usbplug_v1.01.bin

| Date       | File                   | Build commit | Severity  |
| ---------- | :--------------------- | ------------ | --------- |
| 2025-04-14 | rv1126b_usbplug_v1.01.bin | 8532419ba16    | important |

### New

1. Add arch timer 1Ghz support.

------

## rv1126b_bl32_v1.01.bin

| Date       | File                   | Build commit | Severity  |
| ---------- | :--------------------- | ------------ | --------- |
| 2025-04-14 | rv1126b_bl32_v1.01.bin | 0e8447122    | important |

### New

1. Add support for oem otp key cipher.

### Fixed

| Index | Severity  | Update                          | Issue description | Issue source |
| ----- | --------- | ------------------------------- | ----------------- | ------------ |
| 1     | important | fixed FW Enryption Key offset   | -                 | -            |
| 2     | important | correct the clock used by RKRNG | -                 | -            |

------

## rv1126b_bl32_v1.00.bin

| Date       | File                   | Build commit | Severity  |
| ---------- | :--------------------- | ------------ | --------- |
| 2025-04-08 | rv1126b_bl32_v1.00.bin | f63ac68db    | important |

### New

1. Initial version.

------

## rv1126b_bl31_v1.01.elf

| Date       | File                   | Build commit | Severity  |
| ---------- | :--------------------- | ------------ | --------- |
| 2025-04-02 | rv1126b_bl31_v1.01.elf | a14bf51f4    | important |

### New

1. Support system suspend.
2. Support scmi clock.

------

## rv1126b_{ddr,spl,usbplug,bl31}_v1.00.bin

| Date       | File                                     | Build commit                                                 | Severity |
| ---------- | :--------------------------------------- | ------------------------------------------------------------ | -------- |
| 2025-03-31 | rv1126b_{ddr,spl,usbplug,bl31}_v1.00.bin | ddr:cca56bbb07#spl:4d9e803d493#usbplug:4d9e803d493#bl31:a2173dab6 | moderate |

### New

1. Initial version.

------

