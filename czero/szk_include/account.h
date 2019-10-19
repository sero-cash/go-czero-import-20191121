//
// Created by tang zhige on 2019/9/23.
//

#ifndef LIBSUPERZK_ACCOUNT_INCLUDE_H
#define LIBSUPERZK_ACCOUNT_INCLUDE_H

#include "constant.h"

extern void superzk_seed2sk(
    const unsigned char seed[SZK_SEED_WIDTH],
    unsigned char sk[SZK_SK_WIDTH]
);

extern int superzk_sk2tk(
    const unsigned char sk[SZK_SK_WIDTH],
    unsigned char tk[SZK_TK_WIDTH]
);

extern int superzk_tk2pk(
    const unsigned char tk[SZK_TK_WIDTH],
    unsigned char pk[SZK_PK_WIDTH]
);

extern int superzk_pk2pkr(
    const unsigned char pk[SZK_PK_WIDTH],
    const unsigned char r[32],
    unsigned char pkr[SZK_PKr_WIDTH]
);

extern int superzk_hpkr(
    const unsigned char pkr[SZK_PKr_WIDTH],
    unsigned char hpkr[SZK_HPKr_WIDTH]
);

extern int superzk_pk_valid(const unsigned char pk[SZK_PK_WIDTH]);

extern int superzk_pkr_valid(const unsigned char pkr[SZK_PKr_WIDTH]);

extern int superzk_my_pkr(
    const unsigned char tk[SZK_TK_WIDTH],
    const unsigned char pkr[SZK_PKr_WIDTH]
);

extern int superzk_sign_pkr(
    const unsigned char sk[SZK_SK_WIDTH],
    const unsigned char h[32],
    const unsigned char pkr[SZK_PKr_WIDTH],
    unsigned char sign[SZK_SIGN_WIDTH]
);

extern int superzk_verify_pkr(
    const unsigned char h[32],
    const unsigned char sign[SZK_SIGN_WIDTH],
    const unsigned char pkr[SZK_PKr_WIDTH]
);

extern int superzk_gen_key(
    const unsigned char pkr[SZK_PKr_WIDTH],
    const unsigned char rsk[32],
    unsigned char key[32],
    unsigned char rpk[32]
);

extern int superzk_fetch_key(
    const unsigned char pkr[SZK_PKr_WIDTH],
    const unsigned char tk[SZK_TK_WIDTH],
    const unsigned char rpk[32],
    unsigned char key[32],
    unsigned char vskr[32]
);

extern int superzk_gen_zpka(
    const unsigned char pkr[SZK_PKr_WIDTH],
    const unsigned char a[32],
    unsigned char zpka[32]
);

extern int superzk_sign_zpka(
    const unsigned char sk[SZK_SK_WIDTH],
    const unsigned char h[32],
    const unsigned char a[32],
    const unsigned char pkr[SZK_PKr_WIDTH],
    unsigned char sign[SZK_SIGN_WIDTH]
);

extern int superzk_verify_zpka(
    const unsigned char h[32],
    const unsigned char sign[SZK_SIGN_WIDTH],
    const unsigned char zpka[32]
);

extern int superzk_gen_nil(
    const unsigned char tk[SZK_TK_WIDTH],
    const unsigned char root_cm[32],
    const unsigned char pkr[SZK_PKr_WIDTH],
    unsigned char nil[32]
);

extern int superzk_sign_nil(
    const unsigned char tk[SZK_TK_WIDTH],
    const unsigned char h[32],
    const unsigned char root_cm[32],
    const unsigned char pkr[SZK_PKr_WIDTH],
    unsigned char sign[SZK_NSIGN_WIDTH]
);

extern int superzk_verify_nil(
    const unsigned char h[32],
    const unsigned char sign[SZK_NSIGN_WIDTH],
    const unsigned char nil[32],
    const unsigned char root_cm[32],
    const unsigned char pkr[SZK_PKr_WIDTH]
);

extern int superzk_nil2cm(
    const unsigned char tk[SZK_TK_WIDTH],
    const unsigned char nil[32],
    const unsigned char baser[32],
    unsigned char root_cm[32]
);


#endif //LIBSUPERZK_ACCOUNT_H
