import Link from "next/link";

export default function Nav() {
    return (
        <ul className="mt-3">
        <li className = "my-1"><Link className="hover:bg-gray-50 hover:text-black" href="/">Home</Link></li>
        <li className = "my-1"><Link className="hover:bg-gray-50 hover:text-black" href="/products">Products</Link></li>
        <li className = "my-1"><Link className="hover:bg-gray-50 hover:text-black" href="/products/create">Create Products</Link></li>
        </ul>
    );
}