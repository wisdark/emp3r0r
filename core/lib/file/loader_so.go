//go:build linux
// +build linux

package file

// from go-lpe:
// elf2str -file loader.so

const LoaderSO_Data = `QlpoNjFBWSZTWaSI9VMAE4r__________________________9___f33___3_-9_____4BPrd77777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777777774SJEmIEZGgGmQmSnpsRQ9TI1VVP__0qqp__5FVU___U9qTTSqqf_7FVVH_-1VVP_95VVU__1VU___Sqqn_-9VVU__9VVT__9VVT_9VU__Kqn_6qp__vVPU3pVVQ__aKqn_7Iqqf_-pVUf_siqp__6qqn_5VU__9VVT_9VUf_qqn_6qp_-qqf_tFVQ__bUeqqo__NNVVP__Uqqf_lVR__6VVT__VVT_9VUf_-qqp__6qqn_7VVR_-yVVP_9VVD_801VUf_7VVU__Kqn_-qqgf_vUVVH__qVVH_-qqn_5VU__1VU__1VUEppCE0Jppo9CBU9o000ZJVVR__7Roqqo__81VVT__1VVP__Cqqn__sIeoqqp__-gZKqqf_-VVU__9pVVP_8qqn_-VVR__tAqqn_-A0Kqp__gVVT__VVR__iqqf_4qqn_-Kqo__9Mqqp__6VVR__6qqj__VVT__VVT__aqqn__qqqf_-mhVVP_9qqqf_5qqqf_6qqf_-lVUP_9VVH__pVVH_-0VVR__sI1VVAf_5qqqf_7VVU__9VVQP_8qqn_-qqn_-NVVQD__0qqgSJTQmQUzKYNJpkwk0_VGJmSPTKaeiVVU__zVVU__z_VTQ9VVU__2VVU__8xGlVVP__aVVU__9VVR__5VVR__4AVVU__9VVR__5VVT__1VVP__0VVUP__Kqqf_-VVU__9qqqj__1VVB__-lVU__xVVD__VVT__VVT__VVT__VVT__1VVP_81VVH_-Cqqf_4qqn_-BVVD__Kqp__iqqP_8Kqp__qqp__iqqf_6qqf_6qqf_6qqf_5VVP_9VVAAf_6qqf_6qqB__lVUA__wSKIEEaZMkbU9CT0p40jPUxCGNBVVU__9lVVT__9NVVU__9oZVVUf__hMJhVVU___Sqqn__4qqqP__aCqqn__mqqqf__oqqp__7VVVH__mJk0qqo__9qqqf_71VVT__Gqqp__5NKqp__iqqf_4Kqp__6qqn_-VVT__VVR__lVUf_6qqf_6qqH__lVVH__qqqP_9VVH_-qqn_-qqg__yqqf_6qqH_-VVT__VVQP_9VVH_-VVR__iqqf_-qqof_6qqf_5VVA__1VUf_-lVU__9KqoEiiaECEZpkmD1JoMp6pjT0gyDEqqp___qaqqp__5VVT__8JVVT__9NVVU___RlVVT__9NVVU__9VVT__9VVU__9VVR__7Iqqo___TTTQVVU___Eqqo__9VVP_9oVVR__5VVT__ZVVR__lVU__2VVUf_-lVUP_9pVVB__lVU__1VU__9Kqp__iqqA__9NVVQf_7VVUf_5VVP_8qqn_-qqn__qqqf_-jVVU__9Kqof_5VVP_8qqj__1VVP_9VVD__VVQ__wqqn_-FVUH_-Cqqf_5VVBiQgmUYmhkI2T0mp6CbU9Kqqp__-p7U1VVT__2jKqqn__iqqn__sqqqf__qqqj__9NPVVVT__8iqqj__9VVU___Sqqn__lVVP__yVVU__8xVVU___1Kqqf_-qqp__7Kqqn__5VVU__9qqqj__2VVVP__9RVVR__6qqj__9KqqH__vRFVVP__1VVT__2qqqAP__aVVUH__6NVVUP__1VVH_-qqn__pqqqf_6qqf_-VVUH_-qqh__lVUP_9VVP_8VVR__6qqn__oqqn_-0qqn_-qqn__qqqf_6qqf_7VVUP_9VVP_8qqh__6FVU__1VU__1VUA__1VU0vGpta2Cz-0MPyuT9D7nR_F6mdxOL3XteiL9_WF2v99iQv86tj26n7p_CcyiQPW_ecmYoTfe3Bd-cDnoIdlNpLHG81GmzdU2o8s20NNKWdTQs3v0JeAIbG06YRvX2xvJf-3_QXO37m6Yqi7TSSOw6ng00UUUbQo8RQllllu7aHdTd_fwooo2Ut7atVqVKlGq5h4I60888hN3N2eeetXe0uTzz0aq9rFmiiiavhYQsOiijGwlKUpTrWzTTY-JZlKEsUm5mjjjsUUhRut3J9l76t9gCWxLbmr710YG6fO-w6XBIjVbD3eXQtOIsndxxsqNjw_h7G3yMeutOvZ72M3MdBG3nRXs8zwaFXbj5FMx6CrQz9K7dZKhIWWq0rTkIe7poeRtOijtJdHY2V7LtbO10lerMwuy0_bX83BxuP1fW7nn-LxtPdbfT19TV1-s1N5vtz0Opp62tyOPpTs3e6etxc_V3Bst23HQOsbhC4y4AFVpaw1z--fIjv36rQCoYHaaBtAmgK7RTp_U5fH-f5Le18_Yar6b1F3UXmC437h184CQpZubhUXEJFHyb7EgEueuZ7UwgQ-Flzi76GuvD6vkveCAYaUVmZA0nzYEiA_TmR1JoER_XD7fP7_s-5_rd4_a_D9fu_B3Pscvr-twvNy7t7e_P9cZcPi9Ryfe5mb1dmkpcfjoddCg1b6GviEIOzSFf305NZ_1HG8aJDLb8d87ZqxIttFOvAG-rOF18SbU2nGvyZjtJSNEzQf255v8iglGH2jOL598RdivGEVA3TtV95hT0c3W9S9OiuHeL5qT8I3d7UzMvzGY8DQRvE2wsODDd-hAsn3ZbgN-vU02Pjxtt9OPQ9OWLlSS3-LPxuNx9zdzJ_O_9ncTUs-t6t7j2tpOB4vh3PpbvvODyuD9PgbO1wrvj-LVYmMaJPodHx65ggQqWhANpDYCGPTTDBWTJJCTIQhMJCYTJkMJm22NomYltXNBDbBohgVtC72mBSYDaSUidm9jHW-Z3vQlv0TUK0YuFg5qYmeMzHoljBpzQISTGkjXr00Wi7wIGemQnf5-aGdJK45SgghCBJp4MEPhbzTAASSSSSSSSSSSSSSSeHhDWGs2PJBx0WBhJQlN-yyTG0xw0Q5NfkYZjwtFFBCGQmQidQRipgSZmA3TMzCkhIBspd2hueaOa14bQu8aQCX_XymAkmzGxmeSTASYDaTYkJ3kFB5wSYiGHBI1LBRb5UAg0Jbz09nFAl6eigkjDTAy3t7EArsgxTsbG20DaXWRB1Ovg3fl9frxP63WjjWTD4HSc5p5ENno66MRQBOGqy1s53afpvWHNiS6HOwXtihFpgxo5rerCk0lbahv48QxoaCAyUwNBIEA1N98334baI6Yzk3XXXve1l4-GZVNmxDSSIsmPmuqaFYs1aJVK_7NrQqtjtWLdVdqZlv2KLVNurUvKVWaFHH4luGJLicMxP0pbFobMQcIEAEyvATaM04MpBpBZeQxKLAmYmwINNgNtpDYAEgAQy8w-BNmIEAvQIN6rQOWrnzvkdnZ-428wpfXmXSqgVVUNptjdZbtcFeVo8JRzfgNam1x_v5X5-XwNr1eR97q_y0evh2OR3XI2vT4XH5WLrrG96r8O90X2fT52Bw954vfc96Hgmv1Oq1C6U5p6QQJC0VwunYJL-571HE8eKq1yzKQspgJHr6VjKXoEQiFFeS0eqyLzsJIz6kz2o3_J4_8a27y9j9lbS-D31cz7S6rSQHIxIJAfckDR3jG0hsERGiGkIKHaGVK2SRJgA2kB_h3fNQdnoHYoJEZISyYoWOnf82QH7UQANVEHBRBbf2pEJtcvAAtloF_aigAbwJpJXj17I2S8MR0UGz13kUhKfd3SDNcRatDDasgI2IZzNE_9bkTdhrYhM4gudER93t_bAuREQMbG4iBusiJoht6dJNToC08uKgqlckQuRIQLAF-b0ODIpbdODYwTYgHlWMwQRGwQUqsoDGIsnpdgq1epd7ChtzJpBIJJkHueGEUEVMwA9yViM3dOitZd4a-SAXxdiCTCVX5uHeqh5HP3Flzm0tBoPetGebrjLsKkkbDk0ljTO1CZkoz6M9zD8z8fH42oz_V_TXdh4jYi1eswS4XPoZgeCSCD7hIOm8neYOhKCnf5YvCljMPUQCT7ZJIZwmAfG_SzEMkazGVgx3mn3hFF_qiQIc7fN639pxYLWmDZbnYx1GVjcvcbrwtvoL-3YOPfd0XOVJTIEa_gnVRO19vM42h-J9v7GGaNW8QNcTEkgIb4iitRRRJSooooooq0UWtdrl0Mm3qbbhdPg_tKfZ0MjZZI3ll-FZs2cWLFixYsWLFiwGH4T9Hy8P4-VSkuP22fo-Ko17Fv2MyS7cyg8Lkz4bVVXL3XW8TZT82IxczGiU39p2U4klgu6RNMr2hoX4y52ygTqTDxiIo6bAHR8A_Ffmul5X61lfBzi_cn5NH10UYtE__DBNlnlxQ0bVpE99SjQ5KESdiTxZZvfw7CASml0cQBR82yOsRZGE7IrodLe301GrDuyou6eYqJb7O0KzeqM9n7V1e5N0lKWXyvFw1RFi2pK8FlxdWYd6J-GWywaoQPZQ6T2BAG8aOLLc5bDv_1w3qiu9f1jTxmoXyvH6bCqXAt72HyqxC5NVVMZGT8TmY5bGzYKarWdl_xT7P82qLPb1ieRLZ-JtdhGbfT4sc8vozNsTXEb5DeqkiCY3comamYXOq2KSfmehHZ94VBPnVEFPpyJzG7GF7uhtdghTF0qUaEz1LS8DT32STat_ri3KuQr-xe5v1rJ5xr5Z54XtT5fg-gmraUu-zrLE9tmBGl_5yPx-3k_lxNt8dyubrF4OlNNHn8m_pPZ_6GvO_02mO51XDl8PjAgRLPz27B0tUeVcjiYVBm6Gwn6H3yMOpwJo4duZtfBvLs7ap1MUOaX1HTh8yPn-d7_tp7yTPQCJs-iQQBCNg8m0fO-MsfARml6z3az3cMH1ub8txpe2FXC61rF1mDIT9H4xbMZmZiz481l_1NOMpytx-ly4Y5WkEvgn5z4fRLZ_baGJNpXy_EDYxpg2nmoJQazbIOEIxFBtgxobENjGNg_rhAbaQm02MjAQmxQBpsgMQQBptlEt7f4OcHOcQ9EIhAp5eHAMCzEOCBQXLDRNKaXjEkCGUQO1vc733YxeBpZpoml9sePxPSkoZLnfr8rWxd1PKcl9W3esWkTs0EwkIesUOkDESImZF_yPKlI1OkkXDqzaXr7go_jESRKfgbEvopfPUhwDF_b64P9pI9eBhSSZk3cpv_i3H4rFL0oebMO9pkmcn7aQVKUlgl3bCOBYUAbG2GCYpjxvmpFZ1WYBkSSg5Ju5CDfnEBgjhVqHQCWP71QCnl1xTwKRyK-oPf0xjYwVw4hrHFLnONspE_gkMbDJnriFx9bUyT_x1i0qhpH88cLzd9Bqut1TuPMQFJ8CLxFIgsaJSpNkzzXEaNFkQRKCEmhIl57zIpmMz4H9CqQLd2AqsRVNdadJBPX3SQaNaJyDCkdasex0XVpyud126YEHxpD3YywpepIcOYLlZznvO-0s-m6O5jfWaYyJSx7mFkU_ZPWPCepleNJF5TQaWiObP1Y8vovPxrQYgUs1dyYQkovAyTl33IRi0fP6ST33DopVDfGU08ThK7VBe86CbGXFO3BMqfhZdiYdiiRJKu_Gs_45Bh38CGrbOYpCZD7Cvc2G44_QToQ6E6WpJjOLLjGPYwuVjKEcJ1dFiW9nsvmKrMxBWc7IqDeB2sZrMwQmLnr-zxgBMQshjYRiAX7B5Fa7PXU1sNZHinU4y4HV8Aw729gczvovsdrtcz5fJ_P9Pq4n7cDC2Wy_jqf34Ho1b_ca7i831dpzMnzM77-XxfX9aqqqqqqqqqq1atWrW_tXbovjV4HF52193cDR5X08k8vmsDBH2AEO18mAwWrTHM4Y9WRxrXf9PO43X0C-kp6N8y0SgGMCjlMzfIPy4u-LXpeIenubmvR8U39X2_aXnNiYfsTaXz_dM8_2nQqpr9mXtjh0a4-DlsweZrsnSaRIYk0xoY0vcM2za3ODFls73T6RLH42y5yVePXnXMii3urBokrOTpG3pGoQIO1UmoW_yeM416_yV8Z83K7D-bF-eyRh8O1LXXkjwvQ8Hd5nM42e1nh-1kXus8rR7zYcDeeRvd_pMrKytlwNDwsbsc1dY0MfWNtdkRycm3bysrKysrKp27dZ1u3iBmVvZZ45AWJ4BsfxQPzbK917196UTQGEO5VuZ1J7p3BbJ0LIfU-Zkk6e3Y-4LJ0P0cXvFRlWV8Vq1czcqWard21dhXrFGYzMFkh-bs_T2JmMd6rGDRQF-StFl2wZCjfqDU3ofYzLxr6jFmz_M2Zkk88JRRJp2TFJyRslP-9iZ2KXbzskngZqTODYST4GcpyVrJD_0iSReFxrcxr04Uy_gQGtOa5JIxSIY7uXk1KNjqL63GuzUSgk4OAGjaQbCftYCcMJBSqBrJIM0lHwjJ4OSKBl0i1cXMjQApvqY6GmWNE1QKFr321teBfPuzr0IdirZt3tLQxfsrNValySL46zMrkxCzNp2rUcLUo0s8lc0U9LRIO9F-wTZEnXp2K8aFM9mEWwWYz5WL0E86VW1NtSUTzlhTZKxvkGWUY92o214nIoKzAJMfRHPwVZMZYe0k4R3MtZ4r_Ca6YimeEsKxkBGfM9nHDjWrmdtau8xpmNl6HPgRxzJ64x1XkYpPEV-Uy0tmBAJnS6VcgyZHOcNaltjUiWdBYBxxdp01Uc8T5GAzhG4DIwwMPCnwtLBBUjiTldnK5LJqjiTy7Vq7Tt_JFtMLnc-ZE5q7duFBWoK9l7PeMUM95FWE9su1hrCDWOpTCsRtGjzcJHfj79TNnUgRtomsqbXm8S0pW83fSz5i8jkWmG9IIWrcr2-fLKjIL2PmMfKkwKiPqyFaxVmdFGyZ3NrHCavXuxK3PJbPlhbbaabUYKANj7bQTW9vscTYWTqMxVuV2uPGblWeRrezGE-BcVib_2iNsfzp_roqyPQjAlJqoDCstTQlCAD4gaziS54ehpOTPz6RbaRzDpwNqT5DDfDEJsYxyqqzrzcEi4O4CBJkicKSfsWYNIWqd7bpxTJlNCVhhjsUxOMbbyGoYJMaWi8VCYxooYlx7kiZBx3SDCo0kNoERYg5SYkW22gkYmwUWDYlvkyyDQ2kgssOawVdpItYGTaz9adCYIA1XMQ4lEBIEiwgJkIg3E6jyAIcVrqIaNikuDDym22j5sKV6F0gsIIzIQlNhv2BFgcoH5nUQJA0Se3hYpDw0T5ErMYwHZ938lO1x7BQUGJieBxYmqx7l5IstTjAVhrz2ctoRpFJAhZw0jX0AgEjQlOC0LXtoXuEYysRAye40Zwf3waufvorRjpavb3MmPbT7C3Xm6s1CSRjWqocHTqb3q2_wzU4ZoCuXd1DPZ01u058YwRvOIBkQSLCalFUm_SaGAM-G9xBCnUCEEZ-wF6KIZBVI5BRYhX6pGfe3-QXrNJcqvtFUxkWaBZDC7fp9leT2v6Tpztu1bMDufatW7bUteRhWb0CNJNdvNCAqrWRTgkWe6M2BY5b7ms7wh-RdVQ52WFeah-KEqCAoCxajieF6v-dY6iydHNBRLQHM0EKlrbOBG3YCFu3jMj9NeAqGN66-5L18mIb2B02i9vMa9bJn1pLy2RDmhfnZjFlkbcMmshKVnnT0MOnhoEEUpgngIIEgIIEFRBi3IsZx1v78iPDi_0mCKaMyi3ep3Bf3Y2aRYEtIZfBDTLDtuGzNcG-08yU2h66jD1KCkECGYDnwxNE2TgPj83EGusqzIyDKHC6zuB7gsseS-tabd77q9C0BlFfncuKy11nUdPwK0G2Gjyf-32mkVq-sdbCrjWqnJsGdCmaggTEmu82EVEYoqaMZ0kO6kmwIQSCatiZDrzSBtCdJBRu9PFmEUJaQejDCkiIf1dpszwou_L4jsTBrIZn5el43e70vSwOJXwwRbEVXs82VA-HeElqREYiAqM88gZglVX0aYrRCEwGhtUdNaKI0RDwGYYEICBbbGcNGq4iugwRZXcFs4mCxz-iSH4t27zfv1bhusmPj7Tv9Lmc31PL9jH1cgQcng8PwSX0-Vxfl5WDXxLzQnZ1no_H43T6_O_rG0D4uXW8fd5fR_PRZe7u-HrC5Wik9XAtuLBtjer2cIunVsE2h9xAIDG4kxO0frQdvGQ6zIMKClVOkxL0kTqTBgjSYPRZFwLmCTdn45ej3O7pX-6_dxbiAssS71gL634PP63r1_Z7zy8Z1eh5_S6fqcXdOIetg3nmDffCdSFBx226HXKtf51_uU6av2zmIzOJn1rkMPaeb_ev5N34_C-60czb9dw_wuntPL2NjIhxDkiFJzTHXjR2TOkYu-4MLpYirVNH-0E5zuIdK3lIK7FSY7okhoSG1c4ez3PSm9KxUZYZfnm4MRjOuPf6HOxPo3Etnvso_n98LW_BxzG8f3dFgIp0hDUps4yK7sMaBDwgWETMkGDJDNRPRiyJRic_H0rXbIAK_7WAt-HSD5cF1L6rQXmITZfaQQdn3PY6u5_j_lvrdrwfR_Lwtb5PU7vrN5lc_l_9VVVVVVVVVVVVVVVV2_M1hvkzBda5dUCHEEI0GiZLMaSQjXeAfm-_ACD1k-7RnQC85pLteR-T6EdnDytvgdYrSVaKggNhJVjUpU4S0Srae11KcaUtW8hV_mN9Ys2pi1SFaiC_o-sgTRK1M09HW2iu5Uf7ArDBlEA8S1_C_evAk68AD_gAI5461fGknrxVAsbUxQ8YOD-ajfZ8LcTtVyM-3k5-SsNdddgVftYNfHfsF52vQ3snBizdTLkVKsr_VVF-7oYbNHPprdN9vaQ7VMx6RyA4HJEJxmIE-4o-H5t_WeEztwCwTz1yJRTRTv8OHDhw4cOHDhw4cOHK-LLyb5wyc_nM2nz4fqR_gGf0LuSKcKEhSRHqpg==`
