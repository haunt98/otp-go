# otp-go

OTP cli written in go

## Roadmap

- [x] Implement [HOTP](https://datatracker.ietf.org/doc/html/rfc4226)
- [x] Implement [TOTP](https://datatracker.ietf.org/doc/html/rfc6238)
- [ ] [Aegis Vault](https://github.com/beemdevelopment/Aegis/blob/master/docs/vault.md)
- [ ] Cli with basic command
  - [ ] Add, get, list, delete
  - [ ] Support master password
  - [ ] Save encrypted data using sqlite3 in `~/.local/share`
- [ ] Auto paste in clipboard

## Thanks

- [beemdevelopment/Aegis](https://github.com/beemdevelopment/Aegis)
- [rsc/2fa](https://github.com/rsc/2fa)
- [gokyle/hotp](https://github.com/gokyle/hotp)
- [fosskers/totp](https://github.com/fosskers/totp)
- [zalopay-oss/tokeny](https://github.com/zalopay-oss/tokeny)
- [susam/mintotp](https://github.com/susam/mintotp)
- [yitsushi/totp-cli](https://github.com/yitsushi/totp-cli)
- [paolostivanin/OTPClient](https://github.com/paolostivanin/OTPClient)
- [TOTP from scratch](https://codingindex.xyz/2021/03/07/totp-from-scratch/)
- [2FA TOTP without crappy authenticator apps](https://www.codemadness.org/totp.html)
