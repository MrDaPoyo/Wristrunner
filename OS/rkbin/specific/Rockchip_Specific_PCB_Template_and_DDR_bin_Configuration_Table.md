# Rockchip_Specific_PCB_Template_and_DDR_bin_Configuration_Table.md

## 说明

后面章节描述中，提到的**需要WDQS功能的LPDDR4/LPDDR4X颗粒**，是指<https://redmine.rock-chips.com/documents/49>链接里，Rockchip_DDR_Approved_Vendor_List_verX.XX.pdf中带有Application Note [35]或[36]的颗粒型号，以及Rockchip_DDR_Reference_Vendor_List_verX.XX.pdf中带有Application Note [5]或[6]的颗粒型号。

## RK3568

|                  | Detail                                                       |
| ---------------- | ------------------------------------------------------------ |
| **PCB Template** | RK3568_Template_LP4XD200P132SD4_43x28_1600MHz<br>RK3568_Template_LP4XD200P232DD61346_43R0x28R0_1056MHz<br>RK3568_Template_LP4XD320P164SD6H1_46X28_1600MHz<br>RK3568_Template_LP4XeM254P132SS61346_48R6x33R0_2666Mbps<br>上述模板已支持WDQS的版本，且使用**需要WDQS功能的LPDDR4/LPDDR4X颗粒时**，才需要使用下面的DDR BIN |
| **DDR BIN**      | rk3568_ddr_XXXXMHz_wdqs_vX.XX.bin<br/>rk3568_ddr_XXXXMHz_wdqs_D4_LP4_4x_eyescan_vX.XX.bin<br/>rk3568_ddr_XXXXMHz_wdqs_LP4_4x_max_freq_scan_vX.XX.bin<br/>rk3568_ddr_D3_LP3_D4_XXXMHz_LP4_4x_XXXMHz_wdqs_full_space_test_vX.XX.bin |
| **Note**         | 这些模板，必须配合此表格中所列DDR BIN才可支持，需要WDQS功能的LPDDR4/LPDDR4X颗粒。其中<br>rk3568_ddr_XXXXMHz_wdqs_vX.XX.bin，该bin用于生成系统固件<br/>rk3568_ddr_XXXXMHz_wdqs_D4_LP4_4x_eyescan_vX.XX.bin，该bin用于扫描眼图<br/>rk3568_ddr_XXXXMHz_wdqs_LP4_4x_max_freq_scan_vX.XX.bin，该bin用于辅助判断颗粒在PCB上能运行的最高频率<br/>rk3568_ddr_D3_LP3_D4_XXXMHz_LP4_4x_XXXMHz_wdqs_full_space_test_vX.XX.bin，该bin用于辅助判断颗粒是否存在存储单元异常 |

------

## RK3566

| PCB Template                                                 | DDRBIN                                         | Note                                                         |
| ------------------------------------------------------------ | ---------------------------------------------- | ------------------------------------------------------------ |
| RK3566_Template_LP3D178P232DD414_43R0X25R0_<br/>666MHZ_1056MHZ_H1R6_V10_20251125SQJ | rk3566_ddr_2x178ball_lp3_<br/>666MHz_vx.xx.bin | DDRBIN默认666MHz(1333Mbps)，当仅贴1颗LP3时最高可通过ddrbin_tool修改到1056MHz(颗粒也需要支持1056MHz) |

|                  | Detail                                                       |
| ---------------- | ------------------------------------------------------------ |
| **PCB Template** | RK3566_Template_LP4XD200P132SD4_29x43_1333MHz<br>RK3566_Template_LP4XD200P132SD6_39x25_1066MHz<br>RK3566_Template_LP4XD200P132SS8_35x20_1066MHz<br>RK3566_Template_LP4XD200P232DD61346_43R0x28R0_1848Mbps_2112Mbps<br>RK3566_Template_LP4XD320P164SD6H1_41x24_1066MHz<br>RK3566_Template_LP4D342P164SD6H1_44x32_533MHz<br>RK3566_Template_LP4D366P164SD6H1_44R6x27R4_800MHz<br>RK3566_Template_LP4M254P132SS6H1_34x19_800MHz<br>上述模板已支持WDQS的版本，且使用**需要WDQS功能的LPDDR4/LPDDR4X颗粒时**，才需要使用下面的DDR BIN |
| **DDR BIN**      | rk3566_ddr_XXXXMHz_wdqs_vX.XX.bin<br>rk3566_ddr_XXXXMHz_wdqs_D4_LP4_4x_eyescan_vX.XX.bin<br>rk3566_ddr_XXXXMHz_wdqs_LP4_4x_max_freq_scan_vX.XX.bin<br>rk3566_ddr_D3_LP3_D4_XXXMHz_LP4_4x_XXXMHz_wdqs_full_space_test_vX.XX.bin |
| **Note**         | 这些模板，必须配合此表格中所列DDR BIN才可支持，需要WDQS功能的LPDDR4/LPDDR4X颗粒。其中 <br>rk3566_ddr_XXXXMHz_wdqs_vX.XX.bin，该bin用于生成系统固件<br>rk3566_ddr_XXXXMHz_wdqs_D4_LP4_4x_eyescan_vX.XX.bin，该bin用于扫描眼图<br>rk3566_ddr_XXXXMHz_wdqs_LP4_4x_max_freq_scan_vX.XX.bin，该bin用于辅助判断颗粒在PCB上能运行的最高频率 <br>rk3566_ddr_D3_LP3_D4_XXXMHz_LP4_4x_XXXMHz_wdqs_full_space_test_vX.XX.bin，该bin用于辅助判断颗粒是否存在存储单元异常 |

------

## RK3562

| PCB Template                                                 | DDRBIN                                       | Note                                                         |
| ------------------------------------------------------------ | -------------------------------------------- | ------------------------------------------------------------ |
| RK3562_Template_LP3Q296P164SS61346_<br/>38R7x30R3_1066Mbps_H1R6_V10_20251204 | rk3562_ddr_296ball_lp3_<br/>528MHz_vx.xx.bin |                                                              |
| RK3562_Template_LP3D178P232DD414_38R0X25R0<br/>_1332Mbps_2112Mbps_H1R6_V10_20251217SQJ | rk3562_ddr_296ball_lp3_<br/>528MHz_vx.xx.bin | DDRBIN默认528MHz(1056Mbps)，当贴2颗LP3时可通过ddrbin_tool修改到666MHz，贴1颗时最高可修改到1056MHz(颗粒也需要支持1056MHz) |

|                  | Detail                                                       |
| ---------------- | ------------------------------------------------------------ |
| **PCB Template** | RK3562_Template_LP4XD200P132SD4_31R5x26R3_1333MHz<br>RK3562_Template_LP4XD200P132SD6_28R4x22R8_1333MHz<br>RK3562_Template_LP4XD320P164SD6H1_34X24_1333MHz<br>RK3562_Template_LP4D342P164SD6H1_39x26_800MHz<br>RK3562_Template_LP4D366P164SS6H1_27R0X38R0_800MHz<br>上述模板已支持WDQS的版本，且使用**需要WDQS功能的LPDDR4/LPDDR4X颗粒时**，才需要使用下面的DDR BIN |
| **DDR BIN**      | rk3562_ddr_XXXXMHz_wdqs_vX.XX.bin<br>rk3562_ddr_XXXXMHz_wdqs_LP4_4x_eyescan_vX.XX.bin<br>rk3566_ddr_XXXXMHz_wdqs_LP4_4x_max_freq_scan_vX.XX.bin<br>rk3566_ddr_D3_LP3_D4_XXXMHz_LP4_4x_XXXMHz_wdqs_full_space_test_vX.XX.bin |
| **Note**         | 这些模板，必须配合此表格中所列DDR BIN才可支持，需要WDQS功能的LPDDR4/LPDDR4X颗粒。其中<br>rk3562_ddr_XXXXMHz_wdqs_vX.XX.bin，该bin用于生成系统固件 <br>rk3562_ddr_XXXXMHz_wdqs_D4_LP4_4x_eyescan_vX.XX.bin，该bin用于扫描眼图 <br>rk3562_ddr_XXXXMHz_wdqs_LP4_4x_max_freq_scan_vX.XX.bin，该bin用于辅助判断颗粒在PCB上能运行的最高频率<br>rk3562_ddr_D3_LP3_D4_XXXMHz_LP4_4x_XXXMHz_wdqs_full_space_test_vX.XX.bin，该bin用于辅助判断颗粒是否存在存储单元异常 |

------

## RK3528

|                  | Detail                                                       |
| ---------------- | ------------------------------------------------------------ |
| **PCB Template** | RK3528A_Template_LP4xD200P132SD4_30x34_1056MHz<br>RK3528A_Template_LP4XD200P132SD6_20x29_1056MHz<br>上述模板已支持WDQS的版本，且使用**需要WDQS功能的LPDDR4/LPDDR4X颗粒时**，才需要使用下面的DDR BIN |
| **DDR BIN**      | rk3528_ddr_XXXXMHz_wdqs_vX.XX.bin<br>rk3528_ddr_XXXXMHz_wdqs_LP4_4x_eyescan_vX.XX.bin<br>rk3528_ddr_XXXXMHz_wdqs_lpddr4_4x_max_freq_scan_vX.XX.bin<br>rk3528_ddr_XXXMHz_lpddr4_4x_wdqs_full_space_test_vX.XX.bin<br> |
| **Note**         | 这些模板，必须配合此表格中所列DDR BIN才可支持，需要WDQS功能的LPDDR4/LPDDR4X颗粒。其中<br>rk3528_ddr_XXXXMHz_wdqs_vX.XX.bin，该bin用于生成系统固件<br>rk3528_ddr_XXXXMHz_wdqs_LP4_4x_eyescan_vX.XX.bin，该bin用于扫描眼图<br>rk3528_ddr_XXXXMHz_wdqs_lpddr4_4x_max_freq_scan_vX.XX.bin，该bin用于辅助判断颗粒在PCB上能运行的最高频率<br>rk3528_ddr_XXXMHz_lpddr4_4x_wdqs_full_space_test_vX.XX.bin，该bin用于辅助判断颗粒是否存在存储单元异常 |

## RK3326

| PCB Template                                                | DDRBIN                                       | Note                                                         |
| ----------------------------------------------------------- | -------------------------------------------- | ------------------------------------------------------------ |
| RK3326_Template_LP3Q296P132SS414_<br/>V10_20251226_1048_JZY | rk3326_ddr_296ball_lp3_<br/>333MHz_vx.xx.bin | 除了DDRBIN需要特殊替换外，还需要额外修改kernel里px30-dram-default-timing.dtsi 配置。 |

kernel 里额外修改补丁如下：

```c
diff --git a/arch/arm64/boot/dts/rockchip/px30-dram-default-timing.dtsi b/arch/arm64/boot/dts/rockchip/px30-dram-default-timing.dtsi
index 99fb02048c82..c1ac70ece66e 100644
--- a/arch/arm64/boot/dts/rockchip/px30-dram-default-timing.dtsi
+++ b/arch/arm64/boot/dts/rockchip/px30-dram-default-timing.dtsi
@@ -54,10 +54,10 @@ ddr_timing: ddr_timing {
                lpddr3_odt_dis_freq = <400>;
                phy_lpddr3_odt_dis_freq = <400>;
                lpddr3_drv = <LP3_DS_40ohm>;
-               lpddr3_odt = <LP3_ODT_240ohm>;
-               phy_lpddr3_ca_drv = <PHY_DDR4_LPDDR3_2_RON_RTT_34ohm>;
-               phy_lpddr3_ck_drv = <PHY_DDR4_LPDDR3_2_RON_RTT_43ohm>;
-               phy_lpddr3_dq_drv = <PHY_DDR4_LPDDR3_2_RON_RTT_34ohm>;
+               lpddr3_odt = <LP3_ODT_120ohm>;
+               phy_lpddr3_ca_drv = <PHY_DDR4_LPDDR3_2_RON_RTT_48ohm>;
+               phy_lpddr3_ck_drv = <PHY_DDR4_LPDDR3_2_RON_RTT_48ohm>;
+               phy_lpddr3_dq_drv = <PHY_DDR4_LPDDR3_2_RON_RTT_48ohm>;
                phy_lpddr3_odt = <PHY_DDR4_LPDDR3_2_RON_RTT_240ohm>;

                lpddr4_odt_dis_freq = <800>;
```

------
