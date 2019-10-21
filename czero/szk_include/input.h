//
// Created by tang zhige on 2019/10/12.
//

#ifndef LIBSUPERZK_INPUT_INCLUDES_H
#define LIBSUPERZK_INPUT_INCLUDES_H

extern int superzk_prove_input(
    const unsigned char asset_cm_new[32],
    const unsigned char zpka_input[32],
    const unsigned char nil_input[32],
    const unsigned char anchor_input[32],
    const unsigned char asset_cc[32],
    const unsigned char ar_old[32],
    const unsigned char ar_new[32],
    unsigned long index,
    const unsigned char zpkr[32],
    const unsigned char vskr[32],
    const unsigned char baser[32],
    const unsigned char a[32],
    const unsigned char paths[SZK_PATH_DEPTH*32],
    unsigned long pos,
    unsigned char proof[SZK_PROOF_WIDTH]
);

extern int superzk_verify_input(
    const unsigned char proof[SZK_PROOF_WIDTH],
    const unsigned char asset_cm_new[32],
    const unsigned char zpka_input[32],
    const unsigned char nil_input[32],
    const unsigned char anchor_input[32]
);

#endif //LIBSUPERZK_INPUT_H
