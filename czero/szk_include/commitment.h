//
// Created by tang zhige on 2019/9/24.
//

#ifndef LIBSUPERZK_COMMITMENT_H
#define LIBSUPERZK_COMMITMENT_H

#include "constant.h"

extern int superzk_gen_root_cm_p(
    unsigned long index,
    const unsigned char tkn_currency[32],
    const unsigned char tkn_value[32],
    const unsigned char tkt_category[32],
    const unsigned char tkt_value[32],
    const unsigned char ar[SZK_PKr_WIDTH],
    const unsigned char pkr[SZK_PKr_WIDTH],
    unsigned char root_cm[32]
);

extern int superzk_gen_root_cm_c(
    unsigned long index,
    const unsigned char asset_cm[32],
    const unsigned char pkr[SZK_PKr_WIDTH],
    unsigned char root_cm[32]
);

#endif //LIBSUPERZK_COMMITMENT_H
