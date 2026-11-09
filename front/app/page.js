export default function Home() {
  return (
    <main style={{ fontFamily: 'system-ui, sans-serif', padding: '2rem', maxWidth: 640 }}>
      <h1>idp-test-app-multiple</h1>
      <p>Next.js frontend built by the Ruru platform CI and served from the in-cluster private registry.</p>
      <ul>
        <li>api — Go (PostgreSQL)</li>
        <li>front — Next.js (this app)</li>
        <li>worker — Java</li>
      </ul>
    </main>
  )
}
