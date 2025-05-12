'use client';
// I don't really inderstand the use of this code, but from what i could grap for now 
// it allows to refresh the page by calling the route in the server component which normally
// check authentication insted of using thet cache 

import { useEffect } from "react";
import { useRouter } from "next/navigation";


export function SetDynamicRoute() {
    const router = useRouter();

    useEffect(() => {
        router.refresh();
    }, [router]);
    return <></>;
}