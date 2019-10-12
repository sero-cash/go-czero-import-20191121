//
// Created by tang zhige on 2019/10/12.
//

#ifndef LIBSUPERZK_OUTPUT_INCLUDE_H
#define LIBSUPERZK_OUTPUT_INCLUDE_H

extern int superzk_prove_output(
    const unsigned char tkn_currency[32],
    const unsigned char tkn_value[32],
    const unsigned char tkt_category[32],
    const unsigned char tkt_value[32],
    const unsigned char ar[32],
    const unsigned char asset_cm[32],
    unsigned char proof[SZK_PROOF_WIDTH]
);


extern int superzk_verify_output(
    const unsigned char asset_cm[32],
    const unsigned char proof[SZK_PROOF_WIDTH]
);

#endif //LIBSUPERZK_OUTPUT_H
