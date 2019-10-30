//
// Created by tang zhige on 2019/9/27.
//

#ifndef LIBSUPERZK_CZERO_INCLUDES_H
#define LIBSUPERZK_CZERO_INCLUDES_H

#include "constant.h"

extern int czero_tk2pk(
    const unsigned char tk[SZK_TK_WIDTH],
    unsigned char pk[SZK_PK_WIDTH]
);

extern int czero_pk2pkr(
    const unsigned char pk[SZK_PK_WIDTH],
    const unsigned char r[32],
    unsigned char pkr[SZK_PKr_WIDTH]
);

extern int czero_ismy_pkr(
    const unsigned char tk[SZK_TK_WIDTH],
    const unsigned char pkr[SZK_PKr_WIDTH]
);

extern int czero_fetch_key(
    const unsigned char tk[SZK_TK_WIDTH],
    const unsigned char rpk[32],
    unsigned char key[32]
);

extern int czero_dec_einfo(
    const unsigned char key[32],
    char flag,
    const unsigned char einfo[SZK_INFO_WIDTH],
    unsigned char tkn_currency[32],
    unsigned char tkn_value[32],
    unsigned char tkt_category[32],
    unsigned char tkt_value[32],
    unsigned char rsk[32],
    unsigned char memo[64]
);

extern int czero_sign_pkr(
    const unsigned char h[32],
    const unsigned char sk[SZK_PK_WIDTH],
    const unsigned char pkr[SZK_PKr_WIDTH],
    unsigned char s[SZK_SIGN_WIDTH]
);

extern int czero_verify_pkr(
    const unsigned char h[32],
    const unsigned char s[SZK_SIGN_WIDTH],
    const unsigned char pkr[SZK_PKr_WIDTH]
);

extern int czero_gen_nil(
  const unsigned char sk[SZK_SK_WIDTH],
  const unsigned char root_cm[32],
  unsigned char nil[32]
);

extern int czero_gen_trace(
    const unsigned char tk[SZK_TK_WIDTH],
    const unsigned char root_cm[32],
    unsigned char trace[32]
);

extern int czero_sign_nil_ex(
    //---in---
    const unsigned char hash[32],
    const unsigned char sk[SZK_SK_WIDTH],
    const unsigned char pkr[SZK_PKr_WIDTH],
    const unsigned char root_cm[32],
    //---out---
    unsigned char sign_ret[SZK_NSIGN_WIDTH]
);

extern int czero_verify_nil_ex(
    const unsigned char hash[32],
    const unsigned char sign[SZK_NSIGN_WIDTH],
    const unsigned char pkr[SZK_PKr_WIDTH],
    const unsigned char root_cm[32],
    const unsigned char nil[32]
);


extern int czero_gen_out_cm(
    const unsigned char tkn_currency[32],
    const unsigned char tkn_value[32],
    const unsigned char tkt_category[32],
    const unsigned char tkt_value[32],
    const unsigned char memo[SZK_MEMO_WIDTH],
    const unsigned char pkr[SZK_PKr_WIDTH],
    const unsigned char rsk[32],
    unsigned char out_cm[32]
);

extern void czero_gen_root_cm(
    unsigned long index,
    const unsigned char out_cm[32],
    unsigned char cm[32]
);

extern void czero_merkle_combine(
    const unsigned char l[32],
    const unsigned char r[32],
    unsigned char h[32]
);

extern int czero_til2cm(
    const unsigned char tk[SZK_TK_WIDTH],
    const unsigned char til[32],
    unsigned char root_cm[32]
);


#endif //LIBSUPERZK_CZERO_H
