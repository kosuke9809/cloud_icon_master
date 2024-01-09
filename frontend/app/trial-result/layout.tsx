
export default function TrialResultLayout ({
    children
}:{
    children: React.ReactNode
}) {
    return (
        <section>
            <main className="flex flex-1 justify-center">
                {children}
            </main>
        </section>
    )
}