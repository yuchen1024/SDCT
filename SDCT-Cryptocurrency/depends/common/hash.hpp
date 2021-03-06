/****************************************************************************
this hpp implements hash functions 
*****************************************************************************/
#ifndef __HASH__
#define __HASH__

#include "global.hpp"
#include "../sm/sm3hash.hpp"

/* global variables of hash functions */
const size_t HASH_OUTPUT_LEN = 32;  // hash output = 256-bit string


/* 
    case-tailored hash functions based on SM3
*/

/* H(str) mod q: map a string to a big integer */
void SM3Hash_String_to_BN(string &str, BIGNUM *&y)
{ 
    unsigned char hash_output[HASH_OUTPUT_LEN]; 
    unsigned int hash_len;


    char *buffer = const_cast<char *>(str.data()); 
    SM3Hash(reinterpret_cast<unsigned char *>(buffer), str.size(), hash_output, hash_len);

    BN_bin2bn(hash_output, HASH_OUTPUT_LEN, y);
    BN_nnmod(y, y, order, bn_ctx);
}

/* map an EC point to another EC point, used in pp generation */
void SM3Hash_ECP_to_ECP(EC_POINT *&g, EC_POINT *&h)
{
    unsigned char buffer[POINT_LEN];
    unsigned char hash_output[HASH_OUTPUT_LEN]; 
    unsigned int hash_len;

    EC_POINT *ECP_startpoint = EC_POINT_new(group); 
    EC_POINT_copy(ECP_startpoint, g); 

    /* continue the loop until find a point on curve */
    while(true){
        EC_POINT_point2oct(group, ECP_startpoint, POINT_CONVERSION_COMPRESSED, buffer, POINT_LEN, bn_ctx);
        SM3Hash(buffer, POINT_LEN, hash_output, hash_len);
        // set h to be the first EC point sartisfying the following constraint
        if(EC_POINT_oct2point(group, h, hash_output, POINT_LEN, bn_ctx) == 1 
           && EC_POINT_is_on_curve(group, h, bn_ctx) == 1
           && EC_POINT_is_at_infinity(group, h) == 0) break;
        else EC_POINT_add(group, ECP_startpoint, ECP_startpoint, g, bn_ctx); 
    } 
}


/*
    case-tailored hash functions based on SHA256
*/

/* H(str) mod q: map a string to a big integer */
void Hash_String_to_BN(string &str, BIGNUM *&y)
{ 
    unsigned char hash_output[HASH_OUTPUT_LEN]; 
 
    SHA256_CTX sha256_ctx;
    SHA256_Init(&sha256_ctx);

    char *buffer = const_cast<char *>(str.data()); 
    SHA256(reinterpret_cast<unsigned char *>(buffer), str.size(), hash_output);

    BN_bin2bn(hash_output, HASH_OUTPUT_LEN, y);
    BN_nnmod(y, y, order, bn_ctx);
}

/* map an EC point to another EC point, used in pp generation */
void Hash_ECP_to_ECP(EC_POINT *&g, EC_POINT *&h)
{
    unsigned char buffer[POINT_LEN];
    unsigned char hash_output[HASH_OUTPUT_LEN]; 

    EC_POINT *ECP_startpoint = EC_POINT_new(group); 
    EC_POINT_copy(ECP_startpoint, g); 

    SHA256_CTX sha256_ctx;
    SHA256_Init(&sha256_ctx); 
    /* continue the loop until find a point on curve */
    while(true){
        EC_POINT_point2oct(group, ECP_startpoint, POINT_CONVERSION_COMPRESSED, buffer, POINT_LEN, bn_ctx);
        SHA256(buffer, POINT_LEN, hash_output);
        // set h to be the first EC point sartisfying the following constraint
        if(EC_POINT_oct2point(group, h, hash_output, POINT_LEN, bn_ctx) == 1 
           && EC_POINT_is_on_curve(group, h, bn_ctx) == 1
           && EC_POINT_is_at_infinity(group, h) == 0) break;
        else EC_POINT_add(group, ECP_startpoint, ECP_startpoint, g, bn_ctx); 
    } 
}

#endif

