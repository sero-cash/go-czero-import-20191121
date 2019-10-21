//
// Created by tang zhige on 2019/9/23.
//

#ifndef LIBSUPERZK_CONSTANT_INCLUDE_H
#define LIBSUPERZK_CONSTANT_INCLUDE_H

enum {
    SZK_SUCC=0,
    SZK_ERROR=-1,

    SZK_SEED_WIDTH=32,
    SZK_HPKr_WIDTH=20,
    SZK_SK_WIDTH=64,
    SZK_PK_WIDTH=64,
    SZK_TK_WIDTH=64,
    SZK_PKr_WIDTH=96,
    SZK_SIGN_WIDTH=64,
    SZK_NSIGN_WIDTH=96,
    SZK_PATH_DEPTH=29,
    SZK_PROOF_WIDTH=131,
    SZK_MEMO_WIDTH=64,
    SZK_INFO_WIDTH=
            32+ //currency
            32+ //value
            32+ //category
            32+ //value
            64+ //memo
            32, //ar
};

#endif //LIBSUPERZK_CONSTANT_H
