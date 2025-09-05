'use client';
import { useEffect, useState } from 'react';
import React from 'react';

type Event = {
  id: number;
  tenant_id?: string;
  agent_id?: string;
  type: string;
  source: string;
  severity?: string;
  ts: string;
  payload: Record<string, unknown>;
};

export default function EventsPage() {
  const [items, setItems] = useState<Event[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [page, setPage] = useState(0);

  const base = process.env.NEXT_PUBLIC_ADMIN_API_BASE || 'http://localhost:8081';

  async function fetchPage(p = 0) {
    setLoading(true); setError(null);
    try {
      const res = await fetch(`${base}/v1/admin/events?type=&source=&severity=&from=&to=&limit=&offset=${p*25}`);
      if (!res.ok) throw new Error(`HTTP ${res.status}`);
      const data = await res.json();
      setItems(data.items || []);
    } catch (e:any) {
      setError(e.message);
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => { fetchPage(page); }, [page]);

  return (
    console.log('BASE', base),
    <main className="p-6">
      <h1 className="text-2xl font-bold mb-4">Eventos</h1>
      {loading && <div>Carregando…</div>}
      {error && <div className="text-red-600">Erro: {error}</div>}
      {!loading && !error && (
        <>
          <table className="w-full border text-sm">
            <thead className="bg-gray-50">
              <tr>
                <th className="p-2 border">TS</th>
                <th className="p-2 border">Type</th>
                <th className="p-2 border">Source</th>
                <th className="p-2 border">Severity</th>
                <th className="p-2 border">Agent</th>
                <th className="p-2 border">Payload</th>
              </tr>
            </thead>
            <tbody>
              {items.map(ev => (
                <tr key={ev.id} className="hover:bg-gray-50">
                  <td className="p-2 border">{new Date(ev.ts).toLocaleString()}</td>
                  <td className="p-2 border">{ev.type}</td>
                  <td className="p-2 border">{ev.source}</td>
                  <td className="p-2 border">{ev.severity || '-'}</td>
                  <td className="p-2 border">{ev.agent_id || '-'}</td>
                  <td className="p-2 border"><pre className="text-xs overflow-x-auto">{JSON.stringify(ev.payload, null, 0)}</pre></td>
                </tr>
              ))}
            </tbody>
          </table>
          <div className="mt-3 flex gap-2">
            <button className="px-3 py-1 border rounded" disabled={page===0} onClick={()=>setPage(p=>Math.max(0,p-1))}>Anterior</button>
            <button className="px-3 py-1 border rounded" onClick={()=>setPage(p=>p+1)}>Próxima</button>
          </div>
        </>
      )}
    </main>
  );
}

// web/src/app/events/page.tsx
