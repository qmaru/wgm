import { useState } from "react"
import { Toaster } from "sonner"

import Configs from "@/components/Configs"
import Peers from "@/components/Peers"
import Routes from "@/components/Routes"
import Users from "@/components/Users"
import { Tabs } from "@/components/ui"

export default function App() {
  const [tabIndex, setTabIndex] = useState(0)
  const tabs = ["配置", "节点", "用户", "路由"]
  const pages = [<Configs />, <Peers />, <Users />, <Routes />]

  return (
    <main className="relative flex h-screen flex-col overflow-hidden">
      <div className="absolute left-5 top-3 text-sm text-slate-500">
        v{import.meta.env.VITE_APP_VERSION}
      </div>
      <div className="shrink-0 bg-[#f8fafc]/95 backdrop-blur dark:bg-[#111827]/95">
        <Tabs labels={tabs} selected={tabIndex} onChange={setTabIndex} />
      </div>
      <div className="app-scroll min-h-0 flex-1 overflow-y-auto overscroll-contain">
        {pages[tabIndex]}
      </div>
      <Toaster position="bottom-center" duration={2000} richColors />
    </main>
  )
}
