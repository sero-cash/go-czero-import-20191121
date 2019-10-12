//
// Created by tang zhige on 2019/9/24.
//

#ifndef LIBSUPERZK_ASSET_H
#define LIBSUPERZK_ASSET_H

#include "constant.h"

extern int superzk_gen_tkn_base(
    const unsigned char tkn_currency[32],
    unsigned char base[32]
);

extern int superzk_gen_tkt_base(
    const unsigned char tkt_category[32],
    const unsigned char tkt_value[32],
    unsigned char base[32]
);

extern int superzk_gen_asset_cc(
    const unsigned char tkn_currency[32],
    const unsigned char tkn_value[32],
    const unsigned char tkt_category[32],
    const unsigned char tkt_value[32],
    unsigned char cc[32]
);

extern int superzk_gen_asset_cm(
    const unsigned char tkn_currency[32],
    const unsigned char tkn_value[32],
    const unsigned char tkt_category[32],
    const unsigned char tkt_value[32],
    const unsigned char ar[32],
    unsigned char cm[32],
    unsigned char cc[32]
);

extern int superzk_sign_balance(
    //---in---
    int zin_size,
    const unsigned char* zin_acms,
    const unsigned char* zin_ars,
    int zout_size,
    const unsigned char* zout_acms,
    const unsigned char* zout_ars,
    int oin_size,
    const unsigned char* oin_accs,
    int oout_size,
    const unsigned char* oout_accs,
    const unsigned char hash[32],
    //---out---
    unsigned char bsign[SZK_SIGN_WIDTH],
    unsigned char bcr[32]
);

extern int superzk_verify_balance(
    int zin_size,
    const unsigned char* zin_acms,
    int zout_size,
    const unsigned char* zout_acms,
    int oin_size,
    const unsigned char* oin_accs,
    int oout_size,
    const unsigned char* oout_accs,
    const unsigned char hash[32],
    const unsigned char bsign[SZK_SIGN_WIDTH],
    const unsigned char bcr[32]
);


#endif //LIBSUPERZK_ASSET_H
