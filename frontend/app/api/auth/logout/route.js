import { authOptions } from "../[...nextauth]/route";
import { getServerSession } from "next-auth";
import { getIdToken } from "@/utils/sessionTokenAccessor";


export async function GET() {
    const session = await  getServerSession(authOptions);

    if (session) {
        idToken = await getIdToken();
        var url = `${process.env.END_SESSION_URL}?id_token_hint=${idToken}&post_logout_redirect_uri=${encodeURIComponent(process.env.NEXTAUTH_URL)}}`;

        try {
            const resp = await fetch(url, {method: "GET"});
        } catch (err) {
            console.error("[api/auth/logout] error in keycloakSessionLogOut", err);
            return new Response({ status: 500 });
        }

        return new Response({ status: 200 });
    }
}