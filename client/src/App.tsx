import { useEffect, useState } from "react";
import "./index.css";
import { RefreshCcw } from "lucide-react";

export default function App() {
    const [quote, setQuote] = useState("");
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState("");

    const fetchQuote = async () => {
        setLoading(true);
        setError("");
        try {
            const response = await fetch(
                "https://zip-link-production.up.railway.app/random"
            );
            if (!response.ok) {
                throw new Error("Network response was not ok");
            }
            const data = await response.json();
            setQuote(data.quote);
        } catch (error) {
            setError("Failed to fetch quote. Please try again.");
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        fetchQuote();
    }, []);

    return (
        <div className="flex flex-col items-center justify-between min-h-screen bg-stone-950 py-6 px-4">
            <header className="text-center text-gray-400 py-4 font-medium">
                {" "}
                First non-tutorial Golang project. Repo{" "}
                <a
                    href="https://github.com/filippo-fonseca/how-i-met-your-mother-quote-generator"
                    target="_blank"
                    className="text-white underline hover:font-bold"
                >
                    link
                </a>
                .
            </header>
            <div className="flex flex-col items-center">
                <div className="text-center mb-6 flex flex-col items-center">
                    <img
                        src="https://i.ibb.co/XJq2N55/Group-1.png"
                        width={400}
                    />
                    <h1 className="text-xl font-bold mb-2 text-gray-500 -mt-28">
                        <i>How I Met Your Mother</i>
                    </h1>
                    <h1 className="text-4xl font-bold mb-2 text-white">
                        Quote Generator
                    </h1>
                    <h2 className="text-lg text-gray-500 font-medium">
                        Just press the button to get a dose of wisdom from
                        Barney.
                        <br /> Or Ted. Or Marshall. What about Robin or Lily?
                        Who cares, they're all great!
                    </h2>
                </div>
                <div className="flex items-center justify-center bg-white p-6 rounded-lg shadow-lg w-96 h-32 text-center mb-4">
                    {loading && (
                        <p className="text-lg text-gray-500">Loading...</p>
                    )}
                    {error && (
                        <p className="text-md text-red-500 font-medium">
                            {error}
                        </p>
                    )}
                    {quote && <p className="text-md">{quote}</p>}
                </div>
                <button
                    onClick={fetchQuote}
                    className="px-4 py-2 bg-gray-900 text-white rounded-lg hover:opacity-50 transition-colors"
                >
                    <RefreshCcw />
                </button>
            </div>
            <footer className="text-center text-gray-400 mt-6 py-4 font-medium">
                <p>
                    &copy; Made with ❤️ by{" "}
                    <a
                        href="https://github.com/filippo-fonseca"
                        target="_blank"
                        className="text-white underline hover:font-bold"
                    >
                        filippo fonseca
                    </a>
                    . <br /> An experiment in deploying Golang-based backends
                    with React frontends. <br />
                </p>
            </footer>
        </div>
    );
}
