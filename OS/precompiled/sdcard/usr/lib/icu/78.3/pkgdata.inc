GENCCODE_ASSEMBLY_TYPE=-a gcc
SO=so
SOBJ=so
A=a
LIBPREFIX=lib
LIB_EXT_ORDER=.78.3
COMPILE=cc -D_REENTRANT  -DU_HAVE_ELF_H=1  -DU_ALL_IMPLEMENTATION -DU_ATTRIBUTE_DEPRECATED= -march=armv8-a -O2 -pipe -fno-plt -fexceptions         -Wp,-D_FORTIFY_SOURCE=3 -Wformat -Werror=format-security         -fstack-clash-protection -std=c11 -Wall -pedantic -Wshadow -Wpointer-arith -Wmissing-prototypes -Wwrite-strings   -c
LIBFLAGS=-I/usr/include -DPIC -fPIC
GENLIB=cc -march=armv8-a -O2 -pipe -fno-plt -fexceptions         -Wp,-D_FORTIFY_SOURCE=3 -Wformat -Werror=format-security         -fstack-clash-protection -std=c11 -Wall -pedantic -Wshadow -Wpointer-arith -Wmissing-prototypes -Wwrite-strings   -Wl,-O1 -Wl,--sort-common -Wl,--as-needed -Wl,-z,relro -Wl,-z,now  -shared -Wl,-Bsymbolic
LDICUDTFLAGS=
LD_SONAME=-Wl,-soname -Wl,
RPATH_FLAGS=
BIR_LDFLAGS=-Wl,-Bsymbolic
AR=ar
ARFLAGS=r
RANLIB=ranlib
INSTALL_CMD=/usr/bin/install -c
