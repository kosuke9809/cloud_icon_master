import type { Metadata } from 'next'
import { Providers } from './providers'
import Navbar from './components/NavBar'
import './globals.css'
import { Noto_Sans_JP } from 'next/font/google'



const notoSansJp = Noto_Sans_JP({
  weight: ["400", "500"],
  subsets: ["latin"],
  display: "swap",

})

export const metadata: Metadata = {
  title: 'AWS ICON MASTER',
  description: 'aws icon master',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="ja">
      <body className={notoSansJp.className}>
        <Providers>
          <Navbar/>
          {children}
        </Providers>
      </body>
    </html>
  )
}
