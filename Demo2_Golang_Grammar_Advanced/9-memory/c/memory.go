package c

import "C"
import "unsafe"

// Malloc allocates size bytes of memory and returns a pointer to the allocated memory.
func Malloc(size int) unsafe.Pointer {
	return C.malloc(C.size_t(size))
}

// Free frees the memory space pointed to by p, which must have been returned by a previous call to Malloc.
func Free(p unsafe.Pointer) {
	C.free(p)
}

// Memcpy copies n bytes from src to dst. No memory overlap is allowed.
func Memcpy(dst unsafe.Pointer, src []byte, n int) {
	C.memcpy(dst, C.CBytes(src), C.size_t(n))
}

// Memset sets n bytes starting at s to c.
func Memset(s unsafe.Pointer, c int, n int) {
	C.memset(s, C.int(c), C.size_t(n))
}

// Memmove copies n bytes from src to dst. No memory overlap is allowed.
func Memmove(dst, src unsafe.Pointer, length int) {
	C.memmove(dst, src, C.size_t(length))
}
