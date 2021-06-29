#include "../src/SDCT.hpp"

const size_t RANGE_LEN = 32; // set the range to be [0, 2^32-1]
const size_t AGG_NUM = 2; 

void build_test_enviroment()
{
    SplitLine_print('-'); 
    cout << "Build test enviroment for SDCT >>>" << endl; 
    SplitLine_print('-'); 
    cout << "Setup SDCT system" << endl; 
    // setup SDCT system
    size_t SN_LEN = 4;
    size_t IO_THREAD_NUM = 4;      
    size_t DEC_THREAD_NUM = 4; 
    size_t TUNNING = 7; 

    SDCT_SP sp;
    SDCT_SP_new(sp);  

    SDCT_PP pp;
    SDCT_PP_new(pp, RANGE_LEN, AGG_NUM); // allocate memory for pp

    SDCT_Setup(sp, pp, RANGE_LEN, AGG_NUM, SN_LEN, IO_THREAD_NUM, DEC_THREAD_NUM, TUNNING); 

    SDCT_Initialize(pp);

    string SDCT_SP_file = "SDCT.sp"; 
    SDCT_SP_serialize(sp, SDCT_SP_file); 

    string SDCT_PP_file = "SDCT.pp"; 
    SDCT_PP_serialize(pp, SDCT_PP_file); 

    system ("read");

    // create accounts for Alice and Bob
    cout << "Generate two accounts" << endl; 
    SplitLine_print('-'); 

    BIGNUM *alice_balance = BN_new(); 
    BIGNUM *alice_sn = BN_new(); BN_one(alice_sn); 
    BN_set_word(alice_balance, 512);
    SDCT_Account Acct_Alice; 
    SDCT_Account_new(Acct_Alice); 
    SDCT_Create_Account(pp, "Alice", alice_balance, alice_sn, Acct_Alice); 
    string Alice_acct_file = "Alice.account"; 
    SDCT_Account_serialize(Acct_Alice, Alice_acct_file); 
    SDCT_Account_free(Acct_Alice);

    BIGNUM *bob_balance = BN_new(); 
    BIGNUM *bob_sn = BN_new(); BN_one(bob_sn); 
    BN_set_word(bob_balance, 256); 
    SDCT_Account Acct_Bob; 
    SDCT_Account_new(Acct_Bob); 
    SDCT_Create_Account(pp, "Bob", bob_balance, bob_sn, Acct_Bob); 
    string Bob_acct_file = "Bob.account"; 
    SDCT_Account_serialize(Acct_Bob, Bob_acct_file); 
    SDCT_Account_free(Acct_Bob);

    SDCT_SP_free(sp); 
    SDCT_PP_free(pp); 

    system ("read");
} 

void emulate_ctx()
{
    size_t RANGE_LEN = 32; // set the range to be [0, 2^32-1]
    size_t AGG_NUM = 2; 
    
    SDCT_SP sp; 
    SDCT_SP_new(sp); 
    SDCT_SP_deserialize(sp, "SDCT.sp"); 

    SDCT_PP pp; 
    SDCT_PP_new(pp, RANGE_LEN, AGG_NUM); 
    SDCT_PP_deserialize(pp, "SDCT.pp"); 

    SDCT_Account Acct_Alice; 
    SDCT_Account_new(Acct_Alice); 
    SDCT_Account_deserialize(Acct_Alice, "Alice.account"); 
    //SDCT_Account_print(Acct_Alice); 

    SDCT_Account Acct_Bob; 
    SDCT_Account_new(Acct_Bob); 
    SDCT_Account_deserialize(Acct_Bob, "Bob.account"); 
    //SDCT_Account_print(Acct_Bob); 

    cout << "Begin to emulate transactions between Alice and Bob" << endl; 
    SplitLine_print('-'); 
    // cout << "before transactions >>>" << endl; 
    // SplitLine_print('-');
     
    BIGNUM *v = BN_new(); 

    cout << "Wrong Case 1: Invalid CTx --- wrong encryption => equality proof will reject" << endl; 
    SDCT_CTx wrong_ctx1; SDCT_CTx_new(wrong_ctx1);  
    BN_set_word(v, 128); 
    cout << "Alice is going to transfer "<< BN_bn2dec(v) << " to Bob" << endl; 
    SDCT_Create_CTx(pp, Acct_Alice, v, Acct_Bob.pk, wrong_ctx1);

    EC_POINT* noisy = EC_POINT_new(group); 
    ECP_random(noisy); 
    EC_POINT_add(group, wrong_ctx1.transfer.X1, wrong_ctx1.transfer.X1, noisy, bn_ctx);
    EC_POINT_free(noisy); 
    SDCT_Miner(pp, wrong_ctx1, Acct_Alice, Acct_Bob); 
    SDCT_CTx_free(wrong_ctx1); 
    SplitLine_print('-'); 

    system ("read");

    cout << "Wrong Case 2: Invalid CTx --- wrong interval of transfer amount => range proof will reject" << endl; 
    SDCT_CTx wrong_ctx2; SDCT_CTx_new(wrong_ctx2);  
    BN_set_word(v, 4294967296); 
    cout << "Alice is going to transfer "<< BN_bn2dec(v) << " to Bob" << endl; 
    SDCT_Create_CTx(pp, Acct_Alice, v, Acct_Bob.pk, wrong_ctx2);
    SDCT_Miner(pp, wrong_ctx2, Acct_Alice, Acct_Bob); 
    SDCT_CTx_free(wrong_ctx2); 
    SplitLine_print('-'); 

    system ("read");

    cout << "Wrong Case 3: Invalid CTx --- balance is not enough => range proof will reject" << endl; 
    SDCT_CTx wrong_ctx3; SDCT_CTx_new(wrong_ctx3);
    BN_set_word(v, 513);  
    cout << "Alice is going to transfer "<< BN_bn2dec(v) << " coins to Bob" << endl; 
    SDCT_Create_CTx(pp, Acct_Alice, v, Acct_Bob.pk, wrong_ctx3);
    SDCT_Miner(pp, wrong_ctx3, Acct_Alice, Acct_Bob); 
    SDCT_CTx_free(wrong_ctx3); 
    SplitLine_print('-'); 

    system ("read");

    cout << "1st Valid CTx" << endl;
    SDCT_CTx ctx1; SDCT_CTx_new(ctx1);  
    BN_set_word(v, 128); 
    cout << "Alice is going to transfer "<< BN_bn2dec(v) << " coins to Bob" << endl; 
    SDCT_Create_CTx(pp, Acct_Alice, v, Acct_Bob.pk, ctx1);
    SDCT_Miner(pp, ctx1, Acct_Alice, Acct_Bob); 
    SplitLine_print('-'); 

    cout << "After 1st valid transaction >>>>>>" << endl; 
    SplitLine_print('-'); 
    SDCT_Account_print(Acct_Alice); 
    SDCT_Account_print(Acct_Bob); 

    system ("read");

    cout << "2nd Valid CTx" << endl; 
    SDCT_CTx ctx2; SDCT_CTx_new(ctx2);
    BN_set_word(v, 384);  
    cout << "Bob is going to transfer "<< BN_bn2dec(v) << " coins to Alice" << endl; 
    SDCT_Create_CTx(pp, Acct_Bob, v, Acct_Alice.pk, ctx2);
    SDCT_Miner(pp, ctx2, Acct_Bob, Acct_Alice); 
    SplitLine_print('-'); 

    cout << "After 2nd valid transaction >>>>>>" << endl; 
    SplitLine_print('-'); 
    SDCT_Account_print(Acct_Alice); 
    SDCT_Account_print(Acct_Bob); 

    system ("read");

    cout << "Supervision begins >>>" << endl; 
    SplitLine_print('-'); 
    SDCT_Open_CTx(sp, pp, ctx1, v); 
    SplitLine_print('-'); 
    SDCT_Open_CTx(sp, pp, ctx2, v);
    SplitLine_print('-');  
    cout << "Supervision ends >>>" << endl; 
    SplitLine_print('-'); 

    BN_free(v); 
    SDCT_SP_free(sp);
    SDCT_PP_free(pp);
    SDCT_Account_free(Acct_Alice);  
    SDCT_Account_free(Acct_Bob); 

    SDCT_CTx_free(ctx1); 
    SDCT_CTx_free(ctx2); 
}



int main()
{
    // generate the system-wide public parameters   
    global_initialize(NID_X9_62_prime256v1); 
    // setup the system and generate three accounts
    build_test_enviroment(); 
    
    // emulate transactions among Alice, Bob, and Tax office 
    emulate_ctx(); 
    global_finalize(); 

    return 0; 
}



