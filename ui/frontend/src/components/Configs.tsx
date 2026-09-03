import { useEffect, useState } from "react"
import { toast } from "sonner"
import QRCode from "react-qr-code"

import { ClipboardText, DataListAPI } from "@wails/go/backend/App"

import { ApiResponse, DataList, PeerItem } from "@/types"
import { Button } from "@/components/ui"

const steps = ["选择 Interface", "选择 Peers", "确认配置"]

const compareIPv4 = (left: string, right: string) => {
  const leftParts = left.split(".").map(Number)
  const rightParts = right.split(".").map(Number)
  for (let index = 0; index < 4; index += 1) {
    if (leftParts[index] !== rightParts[index]) return leftParts[index] - rightParts[index]
  }
  return 0
}

export default function Configs() {
  const [step, setStep] = useState(0)
  const [peers, setPeers] = useState<PeerItem[]>([])
  const [routes, setRoutes] = useState<DataList["routes"]>([])
  const [iface, setIface] = useState("")
  const [selected, setSelected] = useState<Record<string, boolean>>({})
  const [selectedPeerIds, setSelectedPeerIds] = useState<string[]>([])
  const [routeMap, setRouteMap] = useState<Record<string, string[]>>({})
  const [config, setConfig] = useState("")
  const [copied, setCopied] = useState(false)

  useEffect(() => {
    const load = async () => {
      try {
        const raw = await (DataListAPI as () => Promise<unknown>)()
        const response = raw as ApiResponse<DataList>
        if (response.status === 1) {
          const sortedPeers = [...response.data.peers].sort((a, b) =>
            compareIPv4(a.private_addr, b.private_addr),
          )
          setPeers(sortedPeers)
          setRoutes(response.data.routes)
        }
      } catch {
        toast.error("配置接口请求失败")
      }
    }
    void load()
  }, [])

  const generate = () => {
    if (!iface) return toast.warning("请选择 Interface")
    const chosen = selectedPeerIds
      .map((peerId) => peers.find((peer) => String(peer.id) === peerId))
      .filter((peer): peer is PeerItem => peer !== undefined && String(peer.id) !== iface)

    if (!chosen.length) return toast.warning("请选择 Peers")
    const local = peers.find((peer) => String(peer.id) === iface)

    if (!local) return toast.error("本地接口不存在")
    const interfaceLines = [
      "[Interface]",
      `# ${local.username}`,
      `PrivateKey = ${local.prikey}`,
      `Address = ${local.private_addr}/24`,
      local.port ? `ListenPort = ${local.port}` : "",
      local.mtu ? `MTU = ${local.mtu}` : "",
      local.dns ? `DNS = ${local.dns}` : "",
    ]
      .filter(Boolean)
      .join("\n")

    const peerBlocks = chosen.map((peer) => {
      const ips = [peer.allowed_ips, ...(routeMap[String(peer.id)] || [])]
      return [
        "[Peer]",
        `# ${peer.username}`,
        `PublicKey = ${peer.pubkey}`,
        `AllowedIPs = ${ips.filter(Boolean).join(", ")}`,
        peer.public_addr ? `Endpoint = ${peer.public_addr}:${peer.port}` : "",
        peer.keepalive ? `PersistentKeepalive = ${peer.keepalive}` : "",
      ]
        .filter(Boolean)
        .join("\n")
    })
    setConfig([interfaceLines, ...peerBlocks].join("\n\n"))
    setStep(2)
  }

  const reset = () => {
    setStep(0)
    setIface("")
    setSelected({})
    setSelectedPeerIds([])
    setRouteMap({})
    setConfig("")
    setCopied(false)
  }

  const copy = async () => {
    try {
      const ok = await ClipboardText(config)
      if (ok) {
        setCopied(true)
        toast.success("配置已复制")
      } else toast.error("配置复制失败")
    } catch {
      toast.error("配置复制失败")
    }
  }

  const localPeer = peers.find((peer) => String(peer.id) === iface)
  const chosenPeers = selectedPeerIds
    .map((peerId) => peers.find((peer) => String(peer.id) === peerId))
    .filter((peer): peer is PeerItem => peer !== undefined && String(peer.id) !== iface)
  const stepSummary =
    step === 0 ? "选择 Interface" : step === 1 ? "选择 Peers 和路由" : "配置已生成"

  return (
    <main className="mx-auto min-h-full max-w-6xl px-5 py-6 pb-10 sm:px-8">
      <header className="mb-7 border-b border-slate-200 pb-5 dark:border-slate-700">
        <p className="text-xs font-semibold uppercase tracking-[0.18em] text-wg-accent-strong">
          Configuration builder
        </p>
        <h1 className="mt-1 text-2xl font-semibold tracking-tight text-slate-900 dark:text-white">
          配置
        </h1>
        <p className="mt-1 text-sm text-slate-500">按步骤生成 WireGuard 配置并导出</p>
      </header>
      <div className="mx-auto w-full max-w-4xl">
        <div className="mb-5 grid grid-cols-3 border-b border-slate-200 dark:border-slate-700">
          {steps.map((label, index) => (
            <div
              key={label}
              data-current={index === step}
              data-complete={index < step}
              className="ui-step px-3 py-3 text-center text-sm font-medium"
            >
              <span className="mr-2 font-mono text-xs">0{index + 1}</span>
              {label}
            </div>
          ))}
        </div>
        <div className="ui-panel mb-4 px-4 py-3">
          <p className="text-sm font-medium text-slate-700 dark:text-slate-200">{stepSummary}</p>
          <div className="mt-3 grid gap-1 text-xs text-slate-600 dark:text-slate-300">
            <p>
              <span className="mr-2 text-slate-400">Interface</span>
              {localPeer ? `${localPeer.username} · ${localPeer.private_addr}` : "未选择"}
            </p>
            <p>
              <span className="mr-2 text-slate-400">Peers</span>
              {chosenPeers.length
                ? chosenPeers.map((peer) => `${peer.username} · ${peer.private_addr}`).join("、")
                : "未选择"}
            </p>
          </div>
        </div>
        <div className="mb-6 flex items-center justify-between gap-3 border-b border-slate-200 pb-5 dark:border-slate-700">
          <Button
            variant="outline"
            disabled={step === 0}
            onClick={() => {
              setStep(Math.max(0, step - 1))
            }}
          >
            上一步
          </Button>
          <Button
            onClick={
              step === 1
                ? generate
                : step === 2
                  ? reset
                  : () => {
                      if (!iface) toast.warning("请选择 Interface")
                      else setStep(1)
                    }
            }
          >
            {step === 1 || step === 2 ? "完成" : "下一步"}
          </Button>
        </div>
        {step === 0 && (
          <fieldset>
            <legend className="mb-3 text-sm font-semibold text-slate-800 dark:text-slate-200">
              选择本地接口 Interface
            </legend>
            <div className="grid gap-3">
              {peers.map((peer) => (
                <label
                  key={peer.id}
                  data-selected={iface === String(peer.id)}
                  className="ui-choice-card flex cursor-pointer items-center gap-3 p-4"
                >
                  <input
                    type="radio"
                    name="interface"
                    value={peer.id}
                    checked={iface === String(peer.id)}
                    onChange={(event) => {
                      const nextIface = event.target.value
                      setIface(nextIface)
                      setSelected((current) =>
                        Object.fromEntries(
                          Object.entries(current).filter(([peerId]) => peerId !== nextIface),
                        ),
                      )
                      setSelectedPeerIds((current) =>
                        current.filter((peerId) => peerId !== nextIface),
                      )
                      setRouteMap((current) =>
                        Object.fromEntries(
                          Object.entries(current).filter(([peerId]) => peerId !== nextIface),
                        ),
                      )
                    }}
                  />
                  <span className="min-w-0">
                    <strong
                      className={`block break-all font-mono text-sm font-semibold ${peer.public_addr ? "text-wg-accent-strong" : "text-slate-900 dark:text-white"}`}
                    >
                      {peer.username}
                    </strong>
                    <span className="mt-1 block break-all font-mono text-xs text-slate-500 dark:text-slate-400">
                      {peer.private_addr}
                      {peer.public_addr ? ` · ${peer.public_addr}:${peer.port}` : ""}
                    </span>
                  </span>
                </label>
              ))}
            </div>
          </fieldset>
        )}
        {step === 1 && (
          <fieldset>
            <legend className="mb-3 text-sm font-semibold text-slate-800 dark:text-slate-200">
              选择对端 Peer（可多选）
            </legend>
            <div className="grid gap-3 sm:grid-cols-2">
              {peers
                .filter((peer) => String(peer.id) !== iface)
                .map((peer) => {
                  const peerRoutes = routeMap[String(peer.id)] || []
                  const peerSelected = !!selected[String(peer.id)]
                  const hasPublicAddress = !!peer.public_addr
                  return (
                    <div
                      key={peer.id}
                      data-selected={peerSelected}
                      className="ui-choice-card min-w-0 p-3"
                    >
                      <label className="flex cursor-pointer items-center gap-3">
                        <input
                          type="checkbox"
                          checked={peerSelected}
                          onChange={(event) => {
                            const peerId = String(peer.id)
                            const checked = event.target.checked
                            setSelected((current) => ({ ...current, [peerId]: checked }))
                            setSelectedPeerIds((current) =>
                              checked
                                ? current.includes(peerId)
                                  ? current
                                  : [...current, peerId]
                                : current.filter((id) => id !== peerId),
                            )
                          }}
                        />
                        <span className="min-w-0 flex-1">
                          <strong
                            className={`block break-all font-mono text-sm font-semibold ${hasPublicAddress ? "text-wg-accent-strong" : "text-slate-900 dark:text-white"}`}
                          >
                            {peer.username}
                          </strong>
                          <span className="mt-1 block break-all font-mono text-xs text-slate-500 dark:text-slate-400">
                            {peer.private_addr}
                          </span>
                        </span>
                      </label>
                      <div className="mt-3 flex flex-wrap gap-2 border-t border-dashed border-slate-200 pt-3 dark:border-slate-700">
                        <span className="w-full text-xs text-slate-400">路由（可多选）</span>
                        {routes.length ? (
                          routes.map((route) => {
                            const checked = peerRoutes.includes(route.cidr)
                            return (
                              <Button
                                key={route.id}
                                aria-pressed={checked}
                                disabled={!peerSelected}
                                variant={checked ? "primary" : "outline"}
                                onClick={() => {
                                  setRouteMap((current) => ({
                                    ...current,
                                    [String(peer.id)]: checked
                                      ? peerRoutes.filter((cidr) => cidr !== route.cidr)
                                      : [...peerRoutes, route.cidr],
                                  }))
                                }}
                                className="min-h-0 px-2.5 py-1.5 text-xs font-normal"
                              >
                                {route.cidr}
                              </Button>
                            )
                          })
                        ) : (
                          <span className="text-xs text-slate-400">暂无可选路由</span>
                        )}
                      </div>
                    </div>
                  )
                })}
            </div>
          </fieldset>
        )}
        {step === 2 && (
          <div className="grid gap-4 sm:grid-cols-[1fr_auto]">
            <button
              type="button"
              title="点击复制完整配置"
              onClick={() => void copy()}
              className="ui-preview ui-copyable min-w-0 p-5 text-left"
            >
              <pre className="max-h-[28rem] overflow-auto whitespace-pre-wrap font-mono text-sm leading-6 text-slate-700 dark:text-slate-200">
                {config}
              </pre>
              <p className="mt-4 border-t border-dashed border-slate-200 pt-3 text-xs text-slate-500 dark:border-slate-700">
                点击此处复制完整配置{copied ? " · 已复制" : ""}
              </p>
            </button>
            <div className="ui-preview flex items-center justify-center p-4">
              <QRCode value={config} size={180} />
            </div>
          </div>
        )}
      </div>
    </main>
  )
}
