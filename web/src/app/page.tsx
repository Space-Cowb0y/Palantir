// dashboard
export default function Page(){
  return (
    <main className="p-8">
      <h1 className="text-2xl font-bold">Security Toolkit</h1>
      <div className="grid md:grid-cols-3 gap-4 mt-6">
        <div className="p-4 border rounded-lg">Eventos (24h)</div>
        <div className="p-4 border rounded-lg">Agentes Online</div>
        <div className="p-4 border rounded-lg">Políticas Ativas</div>
      </div>
    </main>
  );
}
