#!/usr/bin/env sh

# !!! This script is a part of distribution packaging system !!!
# !!! Each line of this script was tested and debugged on Linux 32bit, Linux 64bit, FreeBSD !!!
# !!! Change with great care, do not break it !!!

SCHEMA_FILE="org.d2r2.gorsync.gschema.xml"

get_gsettings_schema_file()
{
    local EMBEDDED
    # ***** !!!!! DO NOT REMOVE THIS COMMENT BLOCK - HEREDOC WILL BE POSTED HERE !!!!! *****
    # AUTOMATICALLY_REPLACED_WITH_EMBEDDED_XML_FILE_DECLARATION
    # ***** !!!!! DO NOT REMOVE THIS COMMENT BLOCK - HEREDOC WILL BE POSTED HERE !!!!! *****
    if [ ${#EMBEDDED} -le 0 ]; then
        cat "gsettings/${SCHEMA_FILE}"
    else
        echo "${EMBEDDED}"
    fi
}


# if [ -z  "$1" ]; then
    PREFIX=/usr
    OS_LOWERCASE=$(echo "$OSTYPE" | tr "[:upper:]" "[:lower:]")
    # FreeBSD
    if [ "$OS_LOWERCASE" = "freebsd" ]; then
        PREFIX="${PREFIX}/local"
    # Linux OS
    # elif [[ "$OSTYPE" == "linux-gnu" ]]; then
    # Mac OSX
    # elif [[ "$OSTYPE" == "darwin"* ]]; then
    # POSIX compatibility layer and Linux environment emulation for Windows
    # elif [[ "$OSTYPE" == "cygwin" ]]; then
    # Lightweight shell and GNU utilities compiled for Windows (part of MinGW)
    # elif [[ "$OSTYPE" == "msys" ]]; then
    # Windows
    # elif [[ "$OSTYPE" == "win32" ]]; then
    # else
            # Unknown.
    fi
# else
#    export PREFIX=$1
# fi

if [ "$(id -u)" != "0" ]; then
    # Make sure only root can run our script
    echo "This script must be run as root" 1>&2
    exit 1
fi

# Check availability of required commands
# COMMANDS="install glib-compile-schemas glib-compile-resources msgfmt desktop-file-validate gtk-update-icon-cache"
COMMANDS="install glib-compile-schemas glib-compile-resources msgfmt desktop-file-validate gtk-update-icon-cache"
# if [ "$PREFIX" = '/usr' ] || [ "$PREFIX" = "/usr/local" ]; then
#     COMMANDS="$COMMANDS xdg-desktop-menu"
# fi
# PACKAGES="coreutils glib2 glib2 gettext desktop-file-utils gtk-update-icon-cache xdg-utils"
PACKAGES="coreutils glib2 glib2 gettext desktop-file-utils gtk-update-icon-cache xdg-utils"
i=0
for COMMAND in $COMMANDS; do
    type $COMMAND >/dev/null 2>&1 || {
        j=0
        for PACKAGE in $PACKAGES; do
            if [ $i = $j ]; then
                break
            fi
            j=$(( $j + 1 ))
        done
        echo "Your system is missing command $COMMAND, please install $PACKAGE"
        exit 1
    }
    i=$(( $i + 1 ))
done

SCHEMA_PATH=${PREFIX}/share/glib-2.0/schemas
echo "Installing gsettings schema to ${SCHEMA_PATH}"

# Copy and compile schema
echo "Copying and compiling schema..."
install -d ${SCHEMA_PATH}
# install -m 644 gsettings/${SCHEMA_FILE} ${SCHEMA_PATH}/
echo "$(get_gsettings_schema_file)" > ${SCHEMA_PATH}/${SCHEMA_FILE}
chmod 0644 ${SCHEMA_PATH}/${SCHEMA_FILE}
# Redirect output to /dev/null help on some linux distributions (redhat), which produce
# lot of warnings about "Schema ... are depricated." not related to application.
glib-compile-schemas ${SCHEMA_PATH}/ 2>/dev/null

