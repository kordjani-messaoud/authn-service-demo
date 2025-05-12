import NextAuth from "next-auth";
import KeycloakProvider from "next-auth/providers/keycloak";
import { jwtDecode  } from "jwt-decode";
import { encrypt } from "@/utils/encryption";

// This fuction will refresh access token
async function refreshAccessToken(token) {
    url = `${process.env.REFRECH_TOKEN_URL}`;
    const resp = fetch(url, {
        headers: { "Content-Type": "application/x-www-form-urlencoded", },
        body: new URLSearchParams({
            client_id: process.env.CLIENT_ID,
            client_secret: process.env.CLIENT_SECRET,
            grant_type: "refresh_token",
            refresh_token: token.refresh_token,
        }),
        method: "POST",
    });
    refreshToken = await resp.json();

    if (resp.ok) {
        // copy the previous token object that we previously created with changing some values
        return {
            ...token, // spread operator on the previous token
            access_token: refreshToken.access_token,
            decoded: jwtDecode(refreshToken.access_token),
            id_token: refreshToken.id_token,
            expires_at: Math.floor(Date.now() / 1000) + refreshToken.expires_in,
            refresh_token: refreshToken.refresh_token,
        };
    } else {
        throw new Error("[auth/[...nextauth]] Error in refreshAccessToken: " + refreshToken.json);
    }
}

export const authOptions = {
    providers: [
        KeycloakProvider({
            clientId: `${process.env.CLIENT_ID}`,
            clientSecret: `${process.env.CLIENT_SECRET}`,
            issuer: `${process.env.AUTH_ISSUER}`,
        }),
   ],
   callbacks: {
    async jwt({ token, account }) {
        const nowTimeStamp = Math.floor(Date.now() / 1000);

        if (account) {
            token.decode = jwtDecode(account.access_token);
            token.access_token = account.access_token;
            token.id_token = account.id_token;
            token.expires_at = account.expires_at;
            token.refresh_token = account.refresh_token;
            return token;
        } else if (nowTimeStamp < token.expires_at) {
            // token not expired yet
            return token;
        } else {
            // token expired, needs to refresh
            console.log("Token expired. Will refresh...");
            try {
                const refreshToken = await refreshAccessToken(token);
                console.log("Token refreshed");
                return refreshToken;
            } catch (err) {
                console.error("[auth/[...nextauth]] Error in refreshAccessToken: " + err);
                return { ...token, error: "RefreshAccessTokenError" };
            }
        }
    },
    async session( {session, token }) {

        session.access_token = encrypt(token.access_token);
        session.id_token = encrypt(token.id_token);
        session.roles = token.roles;
        session.error = token.error;
        return session
    }
   }
}

const handler = NextAuth(authOptions);

export { handler as GET, handler as POST};