// I wrote this manually since there aren't a lot of functions exported at the moment.
// would be good to automate this in the future.

int ed25519_secp256k1_proof_size();
char ed25519_secp256k1_prove(char *proof_dst, char *key_dst);
char ed25519_secp256k1_verify(char *src, char *dst);