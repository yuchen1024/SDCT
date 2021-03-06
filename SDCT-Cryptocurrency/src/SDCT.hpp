/****************************************************************************
this hpp implements the SDCT functionality 
*****************************************************************************/
#ifndef __SDCT__
#define __SDCT__

#include "../depends/twisted_elgamal/twisted_elgamal.hpp"        // implement Twisted ElGamal  
#include "../depends/nizk/nizk_plaintext_equality.hpp" // NIZKPoK for plaintext equality
#include "../depends/nizk/nizk_plaintext_knowledge.hpp"        // NIZKPoK for ciphertext/honest encryption 
#include "../depends/nizk/nizk_dlog_equality.hpp"      // NIZKPoK for dlog equality
#include "../depends/bulletproofs/aggregate_bulletproof.hpp"    // implement Log Size Bulletproof

#define DEMO           // demo mode 
//#define DEBUG        // show debug information 

// define the structure of system parameters
struct SDCT_PP{
    size_t RANGE_LEN; // the maximum coin value is 2^RANGE_LEN 
    size_t LOG_RANGE_LEN; // this parameter will be used by Bulletproof
    size_t AGG_NUM;    // number of aggregated proofs (for now, we require m to be the power of 2)
    size_t SN_LEN;    // sn length
    size_t TUNNING; 
    size_t IO_THREAD_NUM; // used by twisted ElGamal 
    size_t DEC_THREAD_NUM; // used by twisted ElGamal

    BIGNUM *BN_MAXIMUM_COINS; 

    EC_POINT *g; 
    EC_POINT *h;
    EC_POINT *u; // used for inside innerproduct statement
    vector<EC_POINT *> vec_g; 
    vector<EC_POINT *> vec_h; // the pp of innerproduct part  

    EC_POINT *pka; // supervisor's pk
    BIGNUM *ska;   // supervisor's sk
};

// define the structure of system parameters
struct SDCT_SP{
    BIGNUM *ska;   // supervisor's sk
};

struct SDCT_Account{
    string identity;     // id
    EC_POINT *pk;              // public key
    BIGNUM *sk;              // secret key
    Twisted_ElGamal_CT balance;  // current balance
    BIGNUM *m;               // dangerous (should only be used for speeding up the proof generation)
    BIGNUM *sn; 
};

// define the structure for confidential transaction
struct SDCT_CTx{
    BIGNUM *sn;                        // serial number: uniquely defines a transaction
    // memo information
    Twisted_ElGamal_CT sender_balance;        // the current balance of pk1 (not necessarily included)
    EC_POINT *pk1; EC_POINT *pk2;      // sender = pk1, receiver = pk2
    Triple_Twisted_ElGamal_CT transfer;    // transfer = (X1 = pk1^r, X2 = pk^2, Y = g^r h^v)
    BIGNUM *v;                         // (defined here only for test, should be remove in the real system)  

    // valid proof
    Plaintext_Equality_Proof plaintext_equality_proof;     // NIZKPoK for transfer ciphertext (X1, X2, Y)
    Bullet_Proof bullet_right_enough_proof;      // aggregated range proof for v and m-v lie in the right range 
    Twisted_ElGamal_CT refresh_sender_updated_balance;  // fresh encryption of updated balance (randomness is known)
    Plaintext_Knowledge_Proof plaintext_knowledge_proof; // NIZKPoK for refresh ciphertext (X^*, Y^*)
    DLOG_Equality_Proof dlog_equality_proof;     // fresh updated balance is correct
};

string Get_ctxfilename(SDCT_CTx &newCTx)
{
    string ctx_file = ECP_ep2string(newCTx.pk1) + "_" + BN_bn2string(newCTx.sn)+".ctx"; 
    return ctx_file; 
}

void SDCT_PP_print(SDCT_PP &pp)
{
    SplitLine_print('-');
    cout << "pp content >>>>>>" << endl; 
    cout << "RANGE_LEN = " << pp.RANGE_LEN << endl; 
    cout << "LOG_RANGE_LEN = " << pp.LOG_RANGE_LEN << endl; 
    cout << "AGG_NUM = " << pp.AGG_NUM << endl; // number of sub-argument (for now, we require m to be the power of 2)

    cout << "SN_LEN = " << pp.SN_LEN << endl;  
    cout << "IO_THREAD_NUM = " << pp.IO_THREAD_NUM << endl; 
    cout << "DEC_THREAD_NUM = " << pp.DEC_THREAD_NUM << endl; 
    cout << "TUNNING = " << pp.TUNNING << endl; 

    ECP_print(pp.g, "g"); 
    ECP_print(pp.h, "h");
    ECP_print(pp.u, "u"); 
    ECP_vec_print(pp.vec_g, "vec_g"); 
    ECP_vec_print(pp.vec_h, "vec_h"); 

    ECP_print(pp.pka, "supervisor's pk"); 
    
    SplitLine_print('-'); 
}

void SDCT_Account_print(SDCT_Account &Acct)
{
    cout << Acct.identity << " account information >>> " << endl;     
    ECP_print(Acct.pk, "pk"); 
    //BN_print(Acct.sk, "sk"); 
    cout << "balance:" << endl; 
    Twisted_ElGamal_CT_print(Acct.balance);  // current balance
    BN_print_dec(Acct.m, "m");  // dangerous (should only be used for speeding up the proof generation)
    BN_print(Acct.sn, "sn"); 
    SplitLine_print('-'); 
}

/* print the details of a confidential transaction */
void SDCT_CTx_print(SDCT_CTx &newCTx)
{
    SplitLine_print('-');
    string ctx_file = Get_ctxfilename(newCTx);  
    cout << ctx_file << " content >>>>>>" << endl; 

    cout << "current sender balance >>>" << endl; 
    Twisted_ElGamal_CT_print(newCTx.sender_balance);
    cout << endl; 

    ECP_print(newCTx.pk1, "pk1"); 
    ECP_print(newCTx.pk2, "pk2"); 
    cout << endl;  

    cout << "transfer >>>" << endl;
    Triple_Twisted_ElGamal_CT_print(newCTx.transfer);
    cout << endl; 

    cout << "NIZKPoK for plaintext equality >>>" << endl; 
    Plaintext_Equality_Proof_print(newCTx.plaintext_equality_proof);
    cout << endl; 

    cout << "refresh updated balance >>>" << endl;
    Twisted_ElGamal_CT_print(newCTx.refresh_sender_updated_balance); 
    cout << endl; 

    cout << "NIZKPoK for refreshing correctness >>>" << endl; 
    DLOG_Equality_Proof_print(newCTx.dlog_equality_proof);
    cout << endl; 

    cout << "NIZKPoK of refresh updated balance >>>" << endl; 
    Plaintext_Knowledge_Proof_print(newCTx.plaintext_knowledge_proof); 
    cout << endl; 

    cout << "range proofs for transfer amount and updated balance >>> " << endl; 
    Bullet_Proof_print(newCTx.bullet_right_enough_proof); 
    cout << endl; 

    SplitLine_print('-'); 
}

void SDCT_SP_new(SDCT_SP &sp)
{
    sp.ska = BN_new(); 
}

void SDCT_SP_free(SDCT_SP &sp)
{
    BN_free(sp.ska); 
}

void SDCT_PP_new(SDCT_PP &pp, size_t RANGE_LEN, size_t AGG_NUM)
{
    pp.BN_MAXIMUM_COINS = BN_new(); 
    pp.g = EC_POINT_new(group); 
    pp.h = EC_POINT_new(group);    
    pp.u = EC_POINT_new(group);

    pp.vec_g.resize(RANGE_LEN*AGG_NUM); ECP_vec_new(pp.vec_g); 
    pp.vec_h.resize(RANGE_LEN*AGG_NUM); ECP_vec_new(pp.vec_h); 

    pp.pka = EC_POINT_new(group); 
    pp.ska = BN_new(); 
}

void SDCT_PP_free(SDCT_PP &pp)
{
    BN_free(pp.BN_MAXIMUM_COINS); 
    EC_POINT_free(pp.g);
    EC_POINT_free(pp.h);
    EC_POINT_free(pp.u); // used for inside innerproduct statement
    ECP_vec_free(pp.vec_g); 
    ECP_vec_free(pp.vec_h); 

    EC_POINT_free(pp.pka); 
    BN_free(pp.ska); 
}

void SDCT_Account_new(SDCT_Account &newAcct){
    newAcct.pk = EC_POINT_new(group);              // public key
    newAcct.sk = BN_new();                         // secret key
    Twisted_ElGamal_CT_new(newAcct.balance);  // current balance
    newAcct.m = BN_new(); 
    newAcct.sn = BN_new(); 
}

void SDCT_Account_free(SDCT_Account &newAcct){
    EC_POINT_free(newAcct.pk);              // public key
    BN_free(newAcct.sk);                         // secret key
    Twisted_ElGamal_CT_free(newAcct.balance);  // current balance
    BN_free(newAcct.m); 
    BN_free(newAcct.sn); 
}

void SDCT_CTx_new(SDCT_CTx &newCTx)
{
    newCTx.sn = BN_new(); 
    newCTx.pk1 = EC_POINT_new(group);
    newCTx.pk2 = EC_POINT_new(group);
    Twisted_ElGamal_CT_new(newCTx.sender_balance); 
    Triple_Twisted_ElGamal_CT_new(newCTx.transfer); 
    newCTx.v = BN_new(); 

    NIZK_Plaintext_Equality_Proof_new(newCTx.plaintext_equality_proof); 
    NIZK_Plaintext_Knowledge_Proof_new(newCTx.plaintext_knowledge_proof); 
    NIZK_DLOG_Equality_Proof_new(newCTx.dlog_equality_proof); 
    Bullet_Proof_new(newCTx.bullet_right_enough_proof); 
    Twisted_ElGamal_CT_new(newCTx.refresh_sender_updated_balance); 
}

void SDCT_CTx_free(SDCT_CTx &newCTx)
{
    BN_free(newCTx.sn); 
    EC_POINT_free(newCTx.pk1);
    EC_POINT_free(newCTx.pk2);
    Twisted_ElGamal_CT_free(newCTx.sender_balance); 
    Triple_Twisted_ElGamal_CT_free(newCTx.transfer); 
    BN_free(newCTx.v); 

    NIZK_Plaintext_Equality_Proof_free(newCTx.plaintext_equality_proof); 
    NIZK_Plaintext_Knowledge_Proof_free(newCTx.plaintext_knowledge_proof); 
    NIZK_DLOG_Equality_Proof_free(newCTx.dlog_equality_proof); 
    Bullet_Proof_free(newCTx.bullet_right_enough_proof); 
    Twisted_ElGamal_CT_free(newCTx.refresh_sender_updated_balance); 
}

// obtain pp for each building block
void Get_Bullet_PP(SDCT_PP &pp, Bullet_PP &bullet_pp)
{
    bullet_pp.RANGE_LEN = pp.RANGE_LEN; 
    bullet_pp.LOG_RANGE_LEN = pp.LOG_RANGE_LEN;
    bullet_pp.AGG_NUM = pp.AGG_NUM;  

    bullet_pp.g = pp.g; 
    bullet_pp.h = pp.h; 
    bullet_pp.u = pp.u; 
    bullet_pp.vec_g = pp.vec_g; 
    bullet_pp.vec_h = pp.vec_h; 
}

void Get_Enc_PP(SDCT_PP &pp, Twisted_ElGamal_PP &enc_pp)
{
    enc_pp.MSG_LEN = pp.RANGE_LEN; 
    enc_pp.TUNNING = pp.TUNNING;
    enc_pp.IO_THREAD_NUM = pp.IO_THREAD_NUM;  
    enc_pp.DEC_THREAD_NUM = pp.DEC_THREAD_NUM;  
    enc_pp.BN_MSG_SIZE = pp.BN_MAXIMUM_COINS; 
    enc_pp.g = pp.g; 
    enc_pp.h = pp.h;  
}

void Get_Plaintext_Equality_PP(SDCT_PP &pp, Plaintext_Equality_PP &pteq_pp)
{
    pteq_pp.g = pp.g; 
    pteq_pp.h = pp.h;  
}

void Get_DLOG_Equality_PP(SDCT_PP &pp, DLOG_Equality_PP &dlogeq_pp)
{
    dlogeq_pp.ss_reserve = "dummy";  
}

void Get_Plaintext_Knowledge_PP(SDCT_PP &pp, Plaintext_Knowledge_PP &ptknowledge_pp)
{
    ptknowledge_pp.g = pp.g; 
    ptknowledge_pp.h = pp.h; 
}

void SDCT_SP_serialize(SDCT_SP &sp, string SDCT_sp_file)
{
    ofstream fout;

    fout.open(SDCT_sp_file, ios::binary); 
    BN_serialize(sp.ska, fout);

    fout.close();   
}

void SDCT_SP_deserialize(SDCT_SP &sp, string SDCT_sp_file)
{
    ifstream fin; 
    fin.open(SDCT_sp_file, ios::binary); 

    BN_deserialize(sp.ska, fin); 

    fin.close();   
}


void SDCT_PP_serialize(SDCT_PP &pp, string SDCT_pp_file)
{
    ofstream fout; 
    fout.open(SDCT_pp_file, ios::binary); 
    fout.write((char *)(&pp.RANGE_LEN), sizeof(pp.RANGE_LEN));
    fout.write((char *)(&pp.LOG_RANGE_LEN), sizeof(pp.LOG_RANGE_LEN));
    fout.write((char *)(&pp.AGG_NUM), sizeof(pp.AGG_NUM));
    fout.write((char *)(&pp.SN_LEN), sizeof(pp.SN_LEN));
    fout.write((char *)(&pp.IO_THREAD_NUM), sizeof(pp.IO_THREAD_NUM));
    fout.write((char *)(&pp.DEC_THREAD_NUM), sizeof(pp.DEC_THREAD_NUM));
    fout.write((char *)(&pp.TUNNING), sizeof(pp.TUNNING));

    BN_serialize(pp.BN_MAXIMUM_COINS, fout);  
    ECP_serialize(pp.g, fout); 
    ECP_serialize(pp.h, fout);
    ECP_serialize(pp.u, fout); 
    ECP_vec_serialize(pp.vec_g, fout); 
    ECP_vec_serialize(pp.vec_h, fout); 

    ECP_serialize(pp.pka, fout); 
    BN_serialize(pp.ska, fout);

    fout.close();   
}

void SDCT_PP_deserialize(SDCT_PP &pp, string SDCT_pp_file)
{
    ifstream fin; 
    fin.open(SDCT_pp_file, ios::binary); 
    fin.read((char *)(&pp.RANGE_LEN), sizeof(pp.RANGE_LEN));
    fin.read((char *)(&pp.LOG_RANGE_LEN), sizeof(pp.LOG_RANGE_LEN));
    fin.read((char *)(&pp.AGG_NUM), sizeof(pp.AGG_NUM));
    fin.read((char *)(&pp.SN_LEN), sizeof(pp.SN_LEN));
    fin.read((char *)(&pp.IO_THREAD_NUM), sizeof(pp.IO_THREAD_NUM));
    fin.read((char *)(&pp.DEC_THREAD_NUM), sizeof(pp.DEC_THREAD_NUM));
    fin.read((char *)(&pp.TUNNING), sizeof(pp.TUNNING));

    BN_deserialize(pp.BN_MAXIMUM_COINS, fin);
    ECP_deserialize(pp.g, fin); 
    ECP_deserialize(pp.h, fin);
    ECP_deserialize(pp.u, fin); 
    ECP_vec_deserialize(pp.vec_g, fin); 
    ECP_vec_deserialize(pp.vec_h, fin); 

    ECP_deserialize(pp.pka, fin); 
    BN_deserialize(pp.ska, fin); 

    fin.close();   
}

void SDCT_Account_serialize(SDCT_Account &user, string SDCT_account_file)
{
    ofstream fout; 
    fout.open(SDCT_account_file, ios::binary);
    fout.write((char *)(&user.identity), sizeof(user.identity));
     
    ECP_serialize(user.pk, fout);              
    BN_serialize(user.sk, fout);             
    Twisted_ElGamal_CT_serialize(user.balance, fout);
    BN_serialize(user.m, fout); 
    BN_serialize(user.sn, fout);
    fout.close();  
};

void SDCT_Account_deserialize(SDCT_Account &user, string SDCT_account_file)
{
    ifstream fin; 
    fin.open(SDCT_account_file, ios::binary);
    fin.read((char *)(&user.identity), sizeof(user.identity));

    ECP_deserialize(user.pk, fin);              
    BN_deserialize(user.sk, fin);             
    Twisted_ElGamal_CT_deserialize(user.balance, fin);
    BN_deserialize(user.m, fin); 
    BN_deserialize(user.sn, fin);
    fin.close();  
};

// save CTx into sn.ctx file
void SDCT_CTx_serialize(SDCT_CTx &newCTx, string SDCT_ctx_file)
{
    ofstream fout; 
    fout.open(SDCT_ctx_file); 
    
    // save sn
    BN_serialize(newCTx.sn, fout); 
     
    // save memo info
    ECP_serialize(newCTx.pk1, fout); 
    ECP_serialize(newCTx.pk2, fout); 
    Triple_Twisted_ElGamal_CT_serialize(newCTx.transfer, fout);
    //BN_serialize(newCTx.v, fout);
    
    // save proofs
    Plaintext_Equality_Proof_serialize(newCTx.plaintext_equality_proof, fout);
    Twisted_ElGamal_CT_serialize(newCTx.refresh_sender_updated_balance, fout); 
    DLOG_Equality_Proof_serialize(newCTx.dlog_equality_proof, fout); 
    Plaintext_Knowledge_Proof_serialize(newCTx.plaintext_knowledge_proof, fout); 
    Bullet_Proof_serialize(newCTx.bullet_right_enough_proof, fout); 
    fout.close();

    // calculate the size of ctx_file
    ifstream fin; 
    fin.open(SDCT_ctx_file, ios::ate | ios::binary);
    cout << SDCT_ctx_file << " size = " << fin.tellg() << " bytes" << endl;
    fin.close(); 
}

/* recover CTx from ctx file */
void SDCT_CTx_deserialize(SDCT_CTx &newCTx, string SDCT_ctx_file)
{
    // Deserialize_CTx(newCTx, ctx_file); 
    ifstream fin; 
    fin.open(SDCT_ctx_file);

    // recover sn
    BN_deserialize(newCTx.sn, fin);
    
    // recover memo
    ECP_deserialize(newCTx.pk1, fin); 
    ECP_deserialize(newCTx.pk2, fin); 
    Triple_Twisted_ElGamal_CT_deserialize(newCTx.transfer, fin);
    //BN_deserialize(newCTx.v, fin); 

    // recover proof
    Plaintext_Equality_Proof_deserialize(newCTx.plaintext_equality_proof, fin);
    Twisted_ElGamal_CT_deserialize(newCTx.refresh_sender_updated_balance, fin); 
    DLOG_Equality_Proof_deserialize(newCTx.dlog_equality_proof, fin); 
    Plaintext_Knowledge_Proof_deserialize(newCTx.plaintext_knowledge_proof, fin); 
    Bullet_Proof_deserialize(newCTx.bullet_right_enough_proof, fin); 
    fin.close(); 
}

/* This function implements Setup algorithm of SDCT */
void SDCT_Setup(SDCT_SP &sp, SDCT_PP &pp, size_t RANGE_LEN, size_t AGG_NUM, 
               size_t SN_LEN, size_t IO_THREAD_NUM, size_t DEC_THREAD_NUM, size_t TUNNING)
{
    pp.RANGE_LEN = RANGE_LEN; 
    pp.LOG_RANGE_LEN = log2(RANGE_LEN); 
    pp.AGG_NUM = AGG_NUM; 
    pp.SN_LEN = SN_LEN;
    pp.IO_THREAD_NUM = IO_THREAD_NUM; 
    pp.DEC_THREAD_NUM = DEC_THREAD_NUM; 
    pp.TUNNING = TUNNING; 

    EC_POINT_copy(pp.g, generator); 
    SM3Hash_ECP_to_ECP(pp.g, pp.h);
    ECP_random(pp.u); // used for inside innerproduct statement
    
    BN_set_word(pp.BN_MAXIMUM_COINS, uint64_t(pow(2, pp.RANGE_LEN)));  

    ECP_vec_random(pp.vec_g); 
    ECP_vec_random(pp.vec_h);

    BN_random(sp.ska); // sk \sample Z_p
    EC_POINT_mul(group, pp.pka, sp.ska, NULL, NULL, bn_ctx); // pka = g^ska  
}

/* initialize the encryption part for faster decryption */
void SDCT_Initialize(SDCT_PP &pp)
{
    cout << "Initialize SDCT >>>" << endl; 
    Twisted_ElGamal_PP enc_pp; 
    Get_Enc_PP(pp, enc_pp);  
    Twisted_ElGamal_Initialize(enc_pp); 
    SplitLine_print('-'); 
}

/* create an account for input identity */
void SDCT_Create_Account(SDCT_PP &pp, string identity, BIGNUM *&init_balance, BIGNUM *&sn, 
                        SDCT_Account &newAcct)
{
    newAcct.identity = identity;
    BN_copy(newAcct.sn, sn);  
    Twisted_ElGamal_PP enc_pp;
    Get_Enc_PP(pp, enc_pp); // enc_pp.g = pp.g, enc_pp.h = pp.h;  

    Twisted_ElGamal_KP keypair; 
    Twisted_ElGamal_KP_new(keypair); 
    Twisted_ElGamal_KeyGen(enc_pp, keypair); // generate a keypair
    EC_POINT_copy(newAcct.pk, keypair.pk); 
    BN_copy(newAcct.sk, keypair.sk);  
    Twisted_ElGamal_KP_free(keypair);  

    BN_copy(newAcct.m, init_balance); 

    // initialize account balance with 0 coins
    BIGNUM *r = BN_new(); 
    SM3Hash_String_to_BN(newAcct.identity, r); 
    Twisted_ElGamal_Enc(enc_pp, newAcct.pk, init_balance, r, newAcct.balance);

    #ifdef DEMO
    cout << identity << "'s account creation succeeds" << endl;
    ECP_print(newAcct.pk, "pk"); 
    cout << identity << "'s initial balance = "; 
    BN_print_dec(init_balance); 
    SplitLine_print('-'); 
    #endif 
}

/* update Account if CTx is valid */
bool SDCT_Update_Account(SDCT_PP &pp, SDCT_CTx &newCTx, SDCT_Account &Acct_sender, SDCT_Account &Acct_receiver)
{    
    Twisted_ElGamal_PP enc_pp; 
    Get_Enc_PP(pp, enc_pp); 
    if (EC_POINT_cmp(group, newCTx.pk1, Acct_sender.pk, bn_ctx) 
        || EC_POINT_cmp(group, newCTx.pk2, Acct_receiver.pk, bn_ctx))
    {
        cout << "sender and receiver addresses do not match" << endl;
        return false;  
    }
    else
    {
        BN_add(Acct_sender.sn, Acct_sender.sn, BN_1);

        Twisted_ElGamal_CT C_out; 
        C_out.X = newCTx.transfer.X1; 
        C_out.Y = newCTx.transfer.Y;
        Twisted_ElGamal_CT C_in; 
        C_in.X = newCTx.transfer.X2; 
        C_in.Y = newCTx.transfer.Y;

        // update sender's balance
        Twisted_ElGamal_HomoSub(Acct_sender.balance, Acct_sender.balance, C_out); 
        // update receiver's balance
        Twisted_ElGamal_HomoAdd(Acct_receiver.balance, Acct_receiver.balance, C_in); 

        Twisted_ElGamal_Dec(enc_pp, Acct_sender.sk, Acct_sender.balance, Acct_sender.m); 
        Twisted_ElGamal_Dec(enc_pp, Acct_receiver.sk, Acct_receiver.balance, Acct_receiver.m);

        SDCT_Account_serialize(Acct_sender, Acct_sender.identity+".account"); 
        SDCT_Account_serialize(Acct_receiver, Acct_receiver.identity+".account"); 
        return true; 
    }
} 

/* reveal the balance */ 
void SDCT_Reveal_Balance(SDCT_PP &pp, SDCT_Account &Acct, BIGNUM *&m)
{
    Twisted_ElGamal_PP enc_pp;
    Get_Enc_PP(pp, enc_pp); 
    Twisted_ElGamal_Dec(enc_pp, Acct.sk, Acct.balance, m); 
    //BN_copy(m, Acct.m); 
}

/* supervisor opens CTx */
void SDCT_Open_CTx(SDCT_SP &sp, SDCT_PP &pp, SDCT_CTx &ctx, BIGNUM *&v)
{
    cout << "Inspect " << Get_ctxfilename(ctx) << endl; 
    auto start_time = chrono::steady_clock::now(); 
    Twisted_ElGamal_PP enc_pp;
    Get_Enc_PP(pp, enc_pp); 

    Twisted_ElGamal_CT CT; 
    CT.X = ctx.transfer.X3;
    CT.Y = ctx.transfer.Y;  
    Twisted_ElGamal_Dec(enc_pp, sp.ska, CT, v); 

    cout << ECP_ep2string(ctx.pk1) << " transfers " << BN_bn2dec(v) 
    << " coins to " << ECP_ep2string(ctx.pk2) << endl; 
    auto end_time = chrono::steady_clock::now(); 
    auto running_time = end_time - start_time;
    cout << "inspecting ctx takes time = " 
    << chrono::duration <double, milli> (running_time).count() << " ms" << endl;
}

/* generate a confidential transaction: pk1 transfers v coins to pk2 */
void SDCT_Create_CTx(SDCT_PP &pp, SDCT_Account &Acct_sender, BIGNUM *&v, EC_POINT *&pk2, SDCT_CTx &newCTx)
{
    #ifdef DEMO
    cout << "begin to genetate CTx >>>>>>" << endl; 
    #endif
    SplitLine_print('-'); 

    auto start_time = chrono::steady_clock::now(); 
    BN_copy(newCTx.sn, Acct_sender.sn);
    EC_POINT_copy(newCTx.pk1, Acct_sender.pk); 
    EC_POINT_copy(newCTx.pk2, pk2); 

    Twisted_ElGamal_PP enc_pp;
    Get_Enc_PP(pp, enc_pp); 

    BIGNUM *r = BN_new(); 
    BN_random(r); 

    BN_copy(newCTx.v, v); 
    Triple_Twisted_ElGamal_Enc(enc_pp, newCTx.pk1, newCTx.pk2, pp.pka, v, r, newCTx.transfer); 

    EC_POINT_copy(newCTx.sender_balance.X, Acct_sender.balance.X);
    EC_POINT_copy(newCTx.sender_balance.Y, Acct_sender.balance.Y);

    #ifdef DEMO
    cout <<"1. generate memo info of CTx" << endl;  
    #endif

    // begin to generate the valid proof for ctx
    string transcript_str = BN_bn2string(newCTx.sn); 

    // generate NIZK proof for validity of transfer              
    Plaintext_Equality_PP pteq_pp; 
    Get_Plaintext_Equality_PP(pp, pteq_pp);
    
    Plaintext_Equality_Instance pteq_instance;
    NIZK_Plaintext_Equality_Instance_new(pteq_instance); 
    EC_POINT_copy(pteq_instance.pk1, newCTx.pk1); 
    EC_POINT_copy(pteq_instance.pk2, newCTx.pk2); 
    EC_POINT_copy(pteq_instance.pk3, pp.pka); 
    EC_POINT_copy(pteq_instance.X1, newCTx.transfer.X1);
    EC_POINT_copy(pteq_instance.X2, newCTx.transfer.X2);
    EC_POINT_copy(pteq_instance.X3, newCTx.transfer.X3);
    EC_POINT_copy(pteq_instance.Y, newCTx.transfer.Y);
    
    Plaintext_Equality_Witness pteq_witness; 
    NIZK_Plaintext_Equality_Witness_new(pteq_witness); 
    BN_copy(pteq_witness.r, r); 
    BN_copy(pteq_witness.v, v); 

    NIZK_Plaintext_Equality_Prove(pteq_pp, pteq_instance, pteq_witness, 
                                  transcript_str, newCTx.plaintext_equality_proof);

    #ifdef DEMO
    cout << "2. generate NIZKPoK for plaintext equality" << endl;  
    #endif

    #ifdef DEMO
    cout << "3. compute updated balance" << endl;  
    #endif
    // compute the updated balance

    Twisted_ElGamal_CT sender_updated_balance; 
    Twisted_ElGamal_CT_new(sender_updated_balance);
    EC_POINT_sub(sender_updated_balance.X, newCTx.sender_balance.X, newCTx.transfer.X1); 
    EC_POINT_sub(sender_updated_balance.Y, newCTx.sender_balance.Y, newCTx.transfer.Y); 

    #ifdef DEMO
    cout << "4. compute refreshed updated balance" << endl;  
    #endif
    // refresh the updated balance (with random coins r^*)
    BIGNUM* r_star = BN_new(); 
    BN_random(r_star);   
    Twisted_ElGamal_ReRand(enc_pp, Acct_sender.pk, Acct_sender.sk, 
                           sender_updated_balance, newCTx.refresh_sender_updated_balance, r_star);

    // generate the NIZK proof for updated balance and fresh updated balance encrypt the same message
    DLOG_Equality_PP dlogeq_pp; 
    Get_DLOG_Equality_PP(pp, dlogeq_pp); 
    
    DLOG_Equality_Instance dlogeq_instance;
    NIZK_DLOG_Equality_Instance_new(dlogeq_instance); 
    // g1 = Y-Y^* = g^{r-r^*}    
    EC_POINT_sub(dlogeq_instance.g1, sender_updated_balance.Y, newCTx.refresh_sender_updated_balance.Y); 
    // h1 = X-X^* = pk^{r-r^*}
    EC_POINT_sub(dlogeq_instance.h1, sender_updated_balance.X, newCTx.refresh_sender_updated_balance.X); 
    
    EC_POINT_copy(dlogeq_instance.g2, enc_pp.g);                         // g2 = g
    EC_POINT_copy(dlogeq_instance.h2, Acct_sender.pk);                    // h2 = pk  
    DLOG_Equality_Witness dlogeq_witness; 
    NIZK_DLOG_Equality_Witness_new(dlogeq_witness); 
    BN_copy(dlogeq_witness.w, Acct_sender.sk); 

    NIZK_DLOG_Equality_Prove(dlogeq_pp, dlogeq_instance, dlogeq_witness, transcript_str, newCTx.dlog_equality_proof); 

    #ifdef DEMO
    cout << "5. generate NIZKPoK for correct refreshing and authenticate the memo info" << endl;  
    #endif

    Plaintext_Knowledge_PP ptke_pp;
    Get_Plaintext_Knowledge_PP(pp, ptke_pp);

    Plaintext_Knowledge_Instance ptke_instance; 
    NIZK_Plaintext_Knowledge_Instance_new(ptke_instance); 
    EC_POINT_copy(ptke_instance.pk, Acct_sender.pk); 
    EC_POINT_copy(ptke_instance.X, newCTx.refresh_sender_updated_balance.X); 
    EC_POINT_copy(ptke_instance.Y, newCTx.refresh_sender_updated_balance.Y); 
    
    Plaintext_Knowledge_Witness ptke_witness; 
    NIZK_Plaintext_Knowledge_Witness_new(ptke_witness);
    BN_copy(ptke_witness.r, r_star); 

    BN_mod_sub(ptke_witness.v, Acct_sender.m, v, order, bn_ctx); // m = Twisted_ElGamal_Dec(sk1, balance); 

    NIZK_Plaintext_Knowledge_Prove(ptke_pp, ptke_instance, ptke_witness, 
                                   transcript_str, newCTx.plaintext_knowledge_proof); 
    
    #ifdef DEMO
    cout << "6. generate NIZKPoK for refreshed updated balance" << endl;  
    #endif

    // aggregated range proof for v and m-v lie in the right range 
    Bullet_PP bullet_pp; 
    Get_Bullet_PP(pp, bullet_pp);
    Bullet_Instance bullet_instance;
    Bullet_Instance_new(bullet_pp, bullet_instance); 
    EC_POINT_copy(bullet_instance.C[0], newCTx.transfer.Y);
    EC_POINT_copy(bullet_instance.C[1], newCTx.refresh_sender_updated_balance.Y);

    Bullet_Witness bullet_witness; 
    Bullet_Witness_new(bullet_pp, bullet_witness); 
    BN_copy(bullet_witness.r[0], pteq_witness.r); 
    BN_copy(bullet_witness.r[1], ptke_witness.r); 
    BN_copy(bullet_witness.v[0], pteq_witness.v);
    BN_copy(bullet_witness.v[1], ptke_witness.v);

    Bullet_Prove(bullet_pp, bullet_instance, bullet_witness, transcript_str, newCTx.bullet_right_enough_proof); 

    #ifdef DEMO
    cout << "7. generate range proofs for transfer amount and updated balance" << endl;    
    #endif

    #ifdef DEMO
    SplitLine_print('-'); 
    #endif

    // #ifndef PREPROCESSING
    // BN_copy(newCTx.v, v); 
    // #endif

    Twisted_ElGamal_CT_free(sender_updated_balance);

    NIZK_Plaintext_Equality_Instance_free(pteq_instance); 
    NIZK_Plaintext_Equality_Witness_free(pteq_witness);

    NIZK_DLOG_Equality_Instance_free(dlogeq_instance);
    NIZK_DLOG_Equality_Witness_free(dlogeq_witness);

    NIZK_Plaintext_Knowledge_Instance_free(ptke_instance); 
    NIZK_Plaintext_Knowledge_Witness_free(ptke_witness); 

    Bullet_Instance_free(bullet_instance); 
    Bullet_Witness_free(bullet_witness);

    auto end_time = chrono::steady_clock::now(); 

    auto running_time = end_time - start_time;
    cout << "ctx generation takes time = " 
    << chrono::duration <double, milli> (running_time).count() << " ms" << endl;
}

/* check if the given confidential transaction is valid */ 
bool SDCT_Verify_CTx(SDCT_PP &pp, SDCT_CTx &newCTx)
{     
    #ifdef DEMO
    cout << "begin to verify CTx >>>>>>" << endl; 
    #endif

    auto start_time = chrono::steady_clock::now(); 
    
    bool Validity; 
    bool V1, V2, V3, V4; 

    string transcript_str = BN_bn2string(newCTx.sn); 

    Plaintext_Equality_PP pteq_pp;
    Get_Plaintext_Equality_PP(pp, pteq_pp); 

    Plaintext_Equality_Instance pteq_instance; 
    NIZK_Plaintext_Equality_Instance_new(pteq_instance); 
    EC_POINT_copy(pteq_instance.pk1, newCTx.pk1);
    EC_POINT_copy(pteq_instance.pk2, newCTx.pk2);
    EC_POINT_copy(pteq_instance.pk3, pp.pka);
    EC_POINT_copy(pteq_instance.X1, newCTx.transfer.X1);
    EC_POINT_copy(pteq_instance.X2, newCTx.transfer.X2);
    EC_POINT_copy(pteq_instance.X3, newCTx.transfer.X3);
    EC_POINT_copy(pteq_instance.Y, newCTx.transfer.Y);

    V1 = NIZK_Plaintext_Equality_Verify(pteq_pp, pteq_instance, transcript_str, newCTx.plaintext_equality_proof);
    
    NIZK_Plaintext_Equality_Instance_free(pteq_instance);

    #ifdef DEMO
    if (V1) cout<< "NIZKPoK for plaintext equality accepts" << endl; 
    else cout<< "NIZKPoK for plaintext equality rejects" << endl; 
    #endif

    // check V2
    Twisted_ElGamal_PP enc_pp; 
    Get_Enc_PP(pp, enc_pp); 

    Twisted_ElGamal_CT updated_sender_balance; 
    Twisted_ElGamal_CT_new(updated_sender_balance);
    EC_POINT_sub(updated_sender_balance.X, newCTx.sender_balance.X, newCTx.transfer.X1); 
    EC_POINT_sub(updated_sender_balance.Y, newCTx.sender_balance.Y, newCTx.transfer.Y); 

    // generate the NIZK proof for updated balance and fresh updated balance encrypt the same message
    DLOG_Equality_PP dlogeq_pp; 
    Get_DLOG_Equality_PP(pp, dlogeq_pp);

    DLOG_Equality_Instance dlogeq_instance; 
    NIZK_DLOG_Equality_Instance_new(dlogeq_instance); 

    EC_POINT_sub(dlogeq_instance.g1, updated_sender_balance.Y, newCTx.refresh_sender_updated_balance.Y); 
    EC_POINT_sub(dlogeq_instance.h1, updated_sender_balance.X, newCTx.refresh_sender_updated_balance.X); 
    EC_POINT_copy(dlogeq_instance.g2, enc_pp.g); 
    EC_POINT_copy(dlogeq_instance.h2, newCTx.pk1);  

    V2 = NIZK_DLOG_Equality_Verify(dlogeq_pp, dlogeq_instance, transcript_str, newCTx.dlog_equality_proof); 

    Twisted_ElGamal_CT_free(updated_sender_balance);
    NIZK_DLOG_Equality_Instance_free(dlogeq_instance); 

    #ifdef DEMO
    if (V2) cout<< "NIZKPoK for refreshing correctness accepts and memo info is authenticated" << endl; 
    else cout<< "NIZKPoK for refreshing correctness rejects or memo info is unauthenticated" << endl; 
    #endif

    Plaintext_Knowledge_PP ptke_pp; 
    Get_Plaintext_Knowledge_PP(pp, ptke_pp);

    Plaintext_Knowledge_Instance ptke_instance; 
    NIZK_Plaintext_Knowledge_Instance_new(ptke_instance); 
    EC_POINT_copy(ptke_instance.pk, newCTx.pk1); 
    EC_POINT_copy(ptke_instance.X, newCTx.refresh_sender_updated_balance.X); 
    EC_POINT_copy(ptke_instance.Y, newCTx.refresh_sender_updated_balance.Y); 

    V3 = NIZK_Plaintext_Knowledge_Verify(ptke_pp, ptke_instance, 
                                         transcript_str, newCTx.plaintext_knowledge_proof); 

    NIZK_Plaintext_Knowledge_Instance_free(ptke_instance); 

    #ifdef DEMO
    if (V3) cout<< "NIZKPoK for refresh updated balance accepts" << endl; 
    else cout<< "NIZKPoK for refresh updated balance rejects" << endl; 
    #endif

    // aggregated range proof for v and m-v lie in the right range 
    Bullet_PP bullet_pp; 
    Get_Bullet_PP(pp, bullet_pp);

    Bullet_Instance bullet_instance;
    Bullet_Instance_new(bullet_pp, bullet_instance); 
    EC_POINT_copy(bullet_instance.C[0], newCTx.transfer.Y);
    EC_POINT_copy(bullet_instance.C[1], newCTx.refresh_sender_updated_balance.Y);

    V4 = Bullet_Verify(bullet_pp, bullet_instance, transcript_str, newCTx.bullet_right_enough_proof); 

    Bullet_Instance_free(bullet_instance); 

    #ifdef DEMO
    if (V4) cout<< "range proofs for transfer amount and updated balance accept" << endl; 
    else cout<< "range proofs for transfer amount and updated balance reject" << endl;   
    #endif

    Validity = V1 && V2 && V3 && V4; 

    string ctx_file  = ECP_ep2string(newCTx.pk1) + "_" + BN_bn2string(newCTx.sn) + ".ctx"; 
    #ifdef DEMO
    if (Validity) cout << ctx_file << " is valid <<<<<<" << endl; 
    else cout << ctx_file << " is invalid <<<<<<" << endl;
    #endif

    auto end_time = chrono::steady_clock::now(); 

    auto running_time = end_time - start_time;
    cout << "ctx verification takes time = " 
    << chrono::duration <double, milli> (running_time).count() << " ms" << endl;

    return Validity; 
}

/* check if a ctx is valid and update accounts if so */
bool SDCT_Miner(SDCT_PP &pp, SDCT_CTx &newCTx, SDCT_Account &Acct_sender, SDCT_Account &Acct_receiver)
{
    if(EC_POINT_cmp(group, newCTx.pk1, Acct_sender.pk, bn_ctx) == 1)
    {
        cout << "sender does not match CTx" << endl; 
        return false; 
    }

    if(EC_POINT_cmp(group, newCTx.pk2, Acct_receiver.pk, bn_ctx) == 1)
    {
        cout << "receiver does not match CTx" << endl; 
        return false; 
    }

    if(SDCT_Verify_CTx(pp, newCTx) == true){
        SDCT_Update_Account(pp, newCTx, Acct_sender, Acct_receiver);
        string ctx_file = ECP_ep2string(newCTx.pk1) + "_" + BN_bn2string(newCTx.sn) + ".ctx"; 
        SDCT_CTx_serialize(newCTx, ctx_file);  
        cout << Get_ctxfilename(newCTx) << " is recorded on the blockchain" << endl; 
        return true; 
    }
    else{
        cout << Get_ctxfilename(newCTx) << " is discarded" << endl; 
        return false; 
    }
}

#endif