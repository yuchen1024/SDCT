/***********************************************************************************
this hpp implements NIZKPoK for three twisited ElGamal ciphertexts 
(randomness reuse) encrypt the same message 
***********************************************************************************/
#ifndef __PTEQ__
#define __PTEQ__

#include "../common/global.hpp"
#include "../common/hash.hpp"
#include "../common/print.hpp"
#include "../common/routines.hpp"

// define structure of PT_EQ_Proof 
struct Plaintext_Equality_PP
{
    EC_POINT *g; 
    EC_POINT *h; 
};

// structure of instance
struct Plaintext_Equality_Instance
{
    EC_POINT *pk1, *pk2, *pk3; 
    EC_POINT *X1, *X2, *X3, *Y; 
};

// structure of witness 
struct Plaintext_Equality_Witness
{
    BIGNUM *v; 
    BIGNUM *r; 
};


// structure of proof 
struct Plaintext_Equality_Proof
{
    EC_POINT *A1, *A2, *A3, *B; // P's first round message
    BIGNUM *z, *t;    // P's response in Zq
};

void NIZK_Plaintext_Equality_Instance_new(Plaintext_Equality_Instance &instance)
{
    instance.pk1 = EC_POINT_new(group);
    instance.pk2 = EC_POINT_new(group);
    instance.pk3 = EC_POINT_new(group);
    instance.X1  = EC_POINT_new(group);
    instance.X2  = EC_POINT_new(group);
    instance.X3  = EC_POINT_new(group);
    instance.Y   = EC_POINT_new(group);
}

void NIZK_Plaintext_Equality_Instance_free(Plaintext_Equality_Instance &instance)
{
    EC_POINT_free(instance.pk1);
    EC_POINT_free(instance.pk2);
    EC_POINT_free(instance.pk3);

    EC_POINT_free(instance.X1);
    EC_POINT_free(instance.X2);
    EC_POINT_free(instance.X3);
    EC_POINT_free(instance.Y);
}

void NIZK_Plaintext_Equality_Witness_new(Plaintext_Equality_Witness &witness)
{
    witness.v = BN_new();
    witness.r = BN_new(); 
}

void NIZK_Plaintext_Equality_Witness_free(Plaintext_Equality_Witness &witness)
{
    BN_free(witness.v);
    BN_free(witness.r); 
}

void NIZK_Plaintext_Equality_Proof_new(Plaintext_Equality_Proof &proof)
{
    proof.A1 = EC_POINT_new(group); 
    proof.A2 = EC_POINT_new(group); 
    proof.A3 = EC_POINT_new(group); 
    proof.B  = EC_POINT_new(group);
    proof.z = BN_new(); 
    proof.t = BN_new();
}

void NIZK_Plaintext_Equality_Proof_free(Plaintext_Equality_Proof &proof)
{
    EC_POINT_free(proof.A1);
    EC_POINT_free(proof.A2);
    EC_POINT_free(proof.A3);
    EC_POINT_free(proof.B);
    BN_free(proof.z);
    BN_free(proof.t);
}


void Plaintext_Equality_Instance_print(Plaintext_Equality_Instance &instance)
{
    cout << "Plaintext Equality Instance >>> " << endl; 
    ECP_print(instance.pk1, "instance.pk1"); 
    ECP_print(instance.pk2, "instance.pk2"); 
    ECP_print(instance.pk3, "instance.pk2"); 
    ECP_print(instance.X1, "instance.X1"); 
    ECP_print(instance.X2, "instance.X2"); 
    ECP_print(instance.X3, "instance.X2"); 
    ECP_print(instance.Y, "instance.Y"); 
} 

void Plaintext_Equality_Witness_print(Plaintext_Equality_Witness &witness)
{
    cout << "Plaintext Equality Witness >>> " << endl; 
    BN_print(witness.v, "witness.v"); 
    BN_print(witness.r, "witness.r"); 
} 

void Plaintext_Equality_Proof_print(Plaintext_Equality_Proof &proof)
{
    SplitLine_print('-'); 
    cout << "NIZKPoK for Plaintext Equality >>> " << endl; 
    ECP_print(proof.A1, "proof.A1"); 
    ECP_print(proof.A2, "proof.A2"); 
    ECP_print(proof.A3, "proof.A3"); 
    ECP_print(proof.B, "proof.B"); 
    BN_print(proof.z, "proof.z"); 
    BN_print(proof.t, "proof.t"); 
} 

void Plaintext_Equality_Proof_serialize(Plaintext_Equality_Proof &proof, ofstream &fout)
{
    ECP_serialize(proof.A1, fout); 
    ECP_serialize(proof.A2, fout);
    ECP_serialize(proof.A3, fout);
    ECP_serialize(proof.B,  fout);
    BN_serialize(proof.z, fout); 
    BN_serialize(proof.t, fout); 
} 

void Plaintext_Equality_Proof_deserialize(Plaintext_Equality_Proof &proof, ifstream &fin)
{
    ECP_deserialize(proof.A1, fin); 
    ECP_deserialize(proof.A2, fin);
    ECP_deserialize(proof.A3, fin);
    ECP_deserialize(proof.B,  fin);
    BN_deserialize(proof.z, fin); 
    BN_deserialize(proof.t, fin); 
} 

/* Setup algorithm */ 
void NIZK_Plaintext_Equality_Setup(Plaintext_Equality_PP &pp)
{ 
    EC_POINT_copy(pp.g, generator); 
    SM3Hash_ECP_to_ECP(pp.g, pp.h);  
}

/* allocate memory for pp */ 
void NIZK_Plaintext_Equality_PP_new(Plaintext_Equality_PP &pp)
{ 
    pp.g = EC_POINT_new(group);
    pp.h = EC_POINT_new(group); 
}

/* free memory of pp */ 
void NIZK_Plaintext_Equality_PP_free(Plaintext_Equality_PP &pp)
{ 
    EC_POINT_free(pp.g); 
    EC_POINT_free(pp.h); 
}

// generate NIZK proof for Ci = Enc(pki, v; r) i={1,2,3} the witness is (r, v)
void NIZK_Plaintext_Equality_Prove(Plaintext_Equality_PP &pp, 
                                   Plaintext_Equality_Instance &instance, 
                                   Plaintext_Equality_Witness &witness, 
                                   string &transcript_str, 
                                   Plaintext_Equality_Proof &proof)
{    
    // initialize the transcript with instance 
    transcript_str += ECP_ep2string(instance.pk1) + ECP_ep2string(instance.pk2) 
                    + ECP_ep2string(instance.pk3) + ECP_ep2string(instance.X1)  
                    + ECP_ep2string(instance.X2)  + ECP_ep2string(instance.X3) 
                    + ECP_ep2string(instance.Y); 

    BIGNUM *a = BN_new(); 
    BIGNUM *b = BN_new(); // the randomness of first round message


    BN_random(a);
    EC_POINT_mul(group, proof.A1, NULL, instance.pk1, a, bn_ctx); // A1 = pk1^a
    EC_POINT_mul(group, proof.A2, NULL, instance.pk2, a, bn_ctx); // A2 = pk2^a
    EC_POINT_mul(group, proof.A3, NULL, instance.pk3, a, bn_ctx); // A3 = pk3^a

    BN_random(b);
    const EC_POINT *vec_A[2]; 
    const BIGNUM *vec_x[2];
    vec_A[0] = pp.g; 
    vec_A[1] = pp.h; 
    vec_x[0] = a; 
    vec_x[1] = b; 
    EC_POINTs_mul(group, proof.B, NULL, 2, vec_A, vec_x, bn_ctx); // B = g^a h^b

    // update the transcript with the first round message
    transcript_str += ECP_ep2string(proof.A1) + ECP_ep2string(proof.A2) 
                    + ECP_ep2string(proof.A3) + ECP_ep2string(proof.B);  
    // compute the challenge
    BIGNUM *e = BN_new(); // V's challenge in Zq
    SM3Hash_String_to_BN(transcript_str, e); // apply FS-transform to generate the challenge

    // compute the response
    BN_mul(proof.z, e, witness.r, bn_ctx); 
    BN_mod_add(proof.z, proof.z, a, order, bn_ctx); // z = a+e*r mod q

    BN_mul(proof.t, e, witness.v, bn_ctx); 
    BN_mod_add(proof.t, proof.t, b, order, bn_ctx); // t = b+e*v mod q

    BN_free(a); 
    BN_free(b);
    BN_free(e);  

    #ifdef DEBUG
    Plaintext_Equality_Proof_print(proof); 
    #endif
}


// check NIZK proof PI for C1 = Enc(pk1, m; r1) and C2 = Enc(pk2, m; r2) the witness is (r1, r2, m)
bool NIZK_Plaintext_Equality_Verify(Plaintext_Equality_PP &pp, 
                                    Plaintext_Equality_Instance &instance, 
                                    string &transcript_str,
                                    Plaintext_Equality_Proof &proof)
{
    // initialize the transcript with instance 
    transcript_str += ECP_ep2string(instance.pk1) + ECP_ep2string(instance.pk2)  
                    + ECP_ep2string(instance.pk3) + ECP_ep2string(instance.X1)   
                    + ECP_ep2string(instance.X2)  + ECP_ep2string(instance.X3)
                    + ECP_ep2string(instance.Y); 

    // update the transcript
    transcript_str += ECP_ep2string(proof.A1) + ECP_ep2string(proof.A2) 
                    + ECP_ep2string(proof.A3) + ECP_ep2string(proof.B);  
    
    // compute the challenge
    BIGNUM *e = BN_new(); 
    SM3Hash_String_to_BN(transcript_str, e); // apply FS-transform to generate the challenge

    bool V1, V2, V3, V4; 
    EC_POINT *LEFT  = EC_POINT_new(group); 
    EC_POINT *RIGHT = EC_POINT_new(group); 
 
    const EC_POINT *vec_A[2]; 
    const BIGNUM *vec_x[2];
    vec_x[0] = BN_1; 
    vec_x[1] = e; 

    // check condition 1
    EC_POINT_mul(group, LEFT, NULL, instance.pk1, proof.z, bn_ctx); // pk1^{z}

    vec_A[0] = proof.A1; 
    vec_A[1] = instance.X1; 
    EC_POINTs_mul(group, RIGHT, NULL, 2, vec_A, vec_x, bn_ctx); 

    V1 = (EC_POINT_cmp(group, LEFT, RIGHT, bn_ctx) == 0); //check pk1^z = A1 X1^e

    // check condition 2
    EC_POINT_mul(group, LEFT, NULL, instance.pk2, proof.z, bn_ctx); // pk2^{z}
    
    vec_A[0] = proof.A2; 
    vec_A[1] = instance.X2; 
    EC_POINTs_mul(group, RIGHT, NULL, 2, vec_A, vec_x, bn_ctx); 

    V2 = (EC_POINT_cmp(group, LEFT, RIGHT, bn_ctx) == 0); //check pk2^z = A2 X2^e

    // check condition 3
    EC_POINT_mul(group, LEFT, NULL, instance.pk3, proof.z, bn_ctx); // pk3^{z}
    
    vec_A[0] = proof.A3; 
    vec_A[1] = instance.X3; 
    EC_POINTs_mul(group, RIGHT, NULL, 2, vec_A, vec_x, bn_ctx); 

    V3 = (EC_POINT_cmp(group, LEFT, RIGHT, bn_ctx) == 0); //check pk3^z = A3 X3^e

    // check condition 4
    vec_A[0] = pp.g; 
    vec_A[1] = pp.h; 
    vec_x[0] = proof.z; 
    vec_x[1] = proof.t; 
    EC_POINTs_mul(group, LEFT, NULL, 2, vec_A, vec_x, bn_ctx); 

    vec_A[0] = proof.B; 
    vec_A[1] = instance.Y; 
    vec_x[0] = BN_1; 
    vec_x[1] = e; 
    EC_POINTs_mul(group, RIGHT, NULL, 2, vec_A, vec_x, bn_ctx); 
    
    V4 = (EC_POINT_cmp(group, LEFT, RIGHT, bn_ctx) == 0); // check g^z h^t = B Y^e

    bool Validity = V1 && V2 && V3 && V4;
    #ifdef DEBUG
    cout << boolalpha << "Condition 1 (Plaintext Equality proof) = " << V1 << endl; 
    cout << boolalpha << "Condition 2 (Plaintext Equality proof) = " << V2 << endl; 
    cout << boolalpha << "Condition 3 (Plaintext Equality proof) = " << V3 << endl; 
    cout << boolalpha << "Condition 4 (Plaintext Equality proof) = " << V4 << endl; 

    if (Validity) 
    { 
        cout<< "NIZK proof for triple twisted ElGamal plaintexts equality accepts >>>" << endl; 
    }
    else 
    {
        cout<< "NIZK proof for triple twisted ElGamal plaintexts equality rejects >>>" << endl; 
    }
    #endif

    BN_free(e); 

    return Validity;
}

#endif



