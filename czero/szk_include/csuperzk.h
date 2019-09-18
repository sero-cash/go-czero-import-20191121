//
// Created by tang zhige on 2019/9/23.
//

#ifndef LIBSUPERZK_CSUPERZK_INCLUDE_H
#define LIBSUPERZK_CSUPERZK_INCLUDE_H

#ifdef __cplusplus
extern "C" {
#endif

#include "account.h"
#include "info.h"
#include "commitment.h"
#include "asset.h"
#include "czero.h"
#include "output.h"
#include "input.h"

extern void superzk_init_params();
extern void superzk_init_params_no_circuit();
extern void superzk_random_pt(unsigned char pt[32]);
extern void superzk_random_fr(unsigned char fr[32]);
extern void superzk_force_fr(unsigned char data[32], unsigned char fr[32]);
extern void superzk_merkle_combine(
    const unsigned char l[32],
    const unsigned char r[32],
    unsigned char h[32]
);

#ifdef __cplusplus
}
#endif

#endif //LIBSUPERZK_CSUPERZK_H
