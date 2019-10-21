//
// Created by tang zhige on 2019/9/24.
//

#ifndef LIBSUPERZK_INFO_INCLUDE_H
#define LIBSUPERZK_INFO_INCLUDE_H

#include "constant.h"

extern int superzk_enc_info(
    const unsigned char key[32],
    const unsigned char tkn_currency[32],
    const unsigned char tkn_value[32],
    const unsigned char tkt_category[32],
    const unsigned char tkt_value[32],
    const unsigned char memo[SZK_MEMO_WIDTH],
    const unsigned char ar[32],
    unsigned char einfo[SZK_INFO_WIDTH]
);

extern int superzk_dec_info(
    const unsigned char key[32],
    const unsigned char einfo[SZK_INFO_WIDTH],
    unsigned char tkn_currency[32],
    unsigned char tkn_value[32],
    unsigned char tkt_category[32],
    unsigned char tkt_value[32],
    unsigned char memo[SZK_MEMO_WIDTH],
    unsigned char ar[32]
);

#endif //LIBSUPERZK_INFO_H
