"use client";

import { useSession, signIn , signOut } from "next-auth/react";
import { useEffect } from "react";

async function keycloakSessionLogOut() {
    try {
        await fetch(`/api/auth/logout`, { method: "GET:" });
    } catch(err) {
        console.error("[authStatus] error in keycloakSessionLogOut", err);
    }
}

export default function AuthStatus() {
    const { data: session, status } = useSession();

    useEffect(() => {
        if (
            status != "loading" &&
            session &&
            session?.error === "RefreshAccessTokenError"
        ){
            signOut({ callbackUrl: "/" });
        }
    }, [session, status]);

    if (status == "loading") {
        return <div className="my-3">Loading...</div>;
    } else if (session) {
        return (
            <div className="my-3">
                logged in as <span className="text-yellow-100">{session.user.email}</span>{" "}
                <button
                    className="bg-blue-900 font-bold text-white py-1 px-2 rounded border border-gray-300 hover:bg-blue-800 hover:border-blue-700 hover:text-white"
                    onclick={() => {
                        keycloakSessionLogOut().then(() => signOut({ callbackUrl: "/" }));
                    }}>
                    Log out
                </button>
            </div>
        )
    }

    return (
        <div className="my-3">
            not logged in . {" "}
            <button
                className="bg-blue-900 font-bold text-white py-1 px-2 rounded border border-gray-300 hover:bg-blue-800 hover:border-blue-700 hover:text-white"
                onClick={() => signIn("keycloak")}>
                Log in
            </button>
        </div>
    )
}