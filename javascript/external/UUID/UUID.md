#  A Universally Unique IDentifier (UUID)

`Request for Comments: 4122`

also known as GUIDs (Globally Unique IDentifier)

A UUID is 128 bits long

The UUID generation algorithm described here supports very high allocation rates of up to 10 million per second per machine if necessary, so that they could even be used as transaction IDs.

## # ABNF

```abnf
UUID                   = time-low "-" time-mid "-"
                        time-high-and-version "-"
                        clock-seq-and-reserved
                        clock-seq-low "-" node
time-low               = 4hexOctet
time-mid               = 2hexOctet
time-high-and-version  = 2hexOctet
clock-seq-and-reserved = hexOctet
clock-seq-low          = hexOctet
node                   = 6hexOctet
hexOctet               = hexDigit hexDigit
hexDigit =
    "0" / "1" / "2" / "3" / "4" / "5" / "6" / "7" / "8" / "9" /
    "a" / "b" / "c" / "d" / "e" / "f" /
    "A" / "B" / "C" / "D" / "E" / "F"

```

`urn:uuid:f81d4fae-7dec-11d0-a765-00a0c91e6bf6`

## # Versions

### - v1

The time-based version

For UUID version 1, the node field consists of an IEEE 802 MAC address, usually the host address.

### - v2

DCE Security version

### - v3

The name-based version `MD5`

### - v4

The randomly or pseudo-randomly generated version

### - v5

The name-based version `SHA-1`