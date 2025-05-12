// Encryption and decryption utilities

import Cryptr from "cryptr";

export function encrypt(text) {
    const secretKey = process.env.NEXTAUTH_SECRET;
    const cryptr = new Cryptr(secretKey);

    const encryptedText = cryptr.encrypt(text);
    return encryptedText;
}


export function decrypt(encryptedText) {
    const secretKey = process.env.NEXTAUTH_SECRET;
    const cryptr = new Cryptr(secretKey);
    const text = cryptr.decrypt(encryptedText);

    return text;
}

