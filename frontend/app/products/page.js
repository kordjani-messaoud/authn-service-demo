import { getServerSession } from "next-auth";
import { authOptions } from "../api/auth/[...nextauth]/route";
import { redirect } from "next/navigation";
import { getAccessToken } from "@/utils/sessionTokenAccessor";
import { SetDynamicRoute } from "@/utils/setDynamicRoute";

// Allow to get all products from backend api
async function getAllProducts() {
    const url = `${process.env.AUTH_SERVICE_URL}/api/v1/products`;
    console.log("BACKEND URL: ", url)

    let accessToken = await getAccessToken();

    const resp = await fetch(url, {
        headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${accessToken}`,
        }
    });

    if (resp.ok) {
        const products = await resp.json();
        return products;
    } else {
        throw new Error("[products] error in getAllProducts: " + resp.status);
    }

}

export default async function Products() {
    const session = await getServerSession(authOptions);
    if (session && session.roles?.includes("viewer")) {
        try {
            const products = await getAllProducts();

            return (
                <main>
                    <SetDynamicRoute></SetDynamicRoute>
                    <h1 className="text-4xl text-center">Products</h1>
                    <table className="border border-gray-500 text-lg ml-auto mr-auto mt-6">
                        <thead>
                            <tr>
                                <th className="bg-blue-900 p-2 border border-gray-500">Id</th>
                                <th className="bg-blue-900 p-2 border border-gray-500">Name</th>
                                <th className="bg-blue-900 p-2 border border-gray-500">
                                    Price
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            {products.map((p) => (
                                <tr key={p.Id}>
                                    <td className="p-1 border border-gray-500">{p.Id}</td>
                                    <td className="p-1 border border-gray-500">{p.Name}</td>
                                    <td className="p-1 border border-gray-500">{p.Price}</td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </main>
            );
        } catch (err) {
            console.error("[products] error in Procunts() because of getAllProducts():", err);

            return (
                <main>
                    <h1 className="text-4xl text-center">Products</h1>
                    <p className="text-red-600 text-center text-lg">
                        Sorry, an error happened. Check the server logs.
                    </p>
                </main>
            );
        }
    } else {
        redirect("/unauthorized");
    }
}