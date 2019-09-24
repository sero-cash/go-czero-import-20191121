//
// Created by tang zhige on 2019/9/23.
//

#ifndef LIBSUPERZK_CSUPERZK_INCLUDE_H
#define LIBSUPERZK_CSUPERZK_INCLUDE_H

#ifdef __cplusplus
extern "C" {
#endif

/*
#include "./light.h"
#include "./keys.h"
#include "./output.h"
#include "./license.h"
#include "./info.h"
#include "./input.h"
#include "./balance.h"
#include "./pkg.h"
#include "./input_s.h"
 */

#include "account.h"
#include "info.h"
#include "commitment.h"
#include "asset.h"

extern void superzk_init_params();
extern void superzk_random_pt(unsigned char pt[32]);
extern void superzk_random_fr(unsigned char fr[32]);
extern void superzk_merkle_combine(
    const unsigned char l[32],
    const unsigned char r[32],
    unsigned char h[32]
);

#ifdef __cplusplus
}
#endif

#endif //LIBSUPERZK_CSUPERZK_H
