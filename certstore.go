package goczmq

/*
#include "czmq.h"
*/
import "C"

// CertStore works with directories of CURVE security certificates.
// It lets you easily load stores from disk and check if a key
// is present or not. This could be done fairly easily in pure
// Go, but is included for the sake of compatibility.
type CertStore struct {
	zcertstoreT *C.struct__zcertstore_t
}

// NewCertStore creates a new certificate store from
// a disk directory, loading and indexing all certificates.
func NewCertStore(location string) *CertStore {
	return &CertStore{
		zcertstoreT: C.zcertstore_new(C.CString(location)),
	}
}

// Insert inserts a certificate into the store in memory.
// Call Save directly on the cert if you wish to save it
// to disk.
func (c *CertStore) Insert(cert *Cert) {
	C.zcertstore_insert(c.zcertstoreT, &cert.zcertT)
}

// Lookup looks up a certificate in the store by public key and
// returns it.
func (c *CertStore) Lookup(key string) *Cert {
	ptr := C.zcertstore_lookup(c.zcertstoreT, C.CString(key))
	if ptr == nil {
		return nil
	}
	return &Cert{
		zcertT: ptr,
	}
}

// Print prints a list of certificates in the store to stdout
func (c *CertStore) Print() {
	C.zcertstore_print(c.zcertstoreT)
}

// Destroy destroys Cert instance
func (c *CertStore) Destroy() {
	C.zcertstore_destroy(&c.zcertstoreT)
}