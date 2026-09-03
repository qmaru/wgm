import { useCallback, useEffect, useState } from "react"
import { toast } from "sonner"

import {
  PeerAddAPI,
  PeerDeleteAPI,
  PeerListAPI,
  PeerUpdateAPI,
  UserListAPI,
} from "@wails/go/backend/App"
import { models } from "@wails/go/models"

import { PeerField, PeerFieldKey, PeerFormState, PeerItem, UserItem, ApiResponse } from "@/types"
import { Button, Card, Dialog, Input, SelectField } from "@/components/ui"

const empty: PeerFormState = {
  user_id: "",
  username: "",
  private_addr: "",
  public_addr: "",
  port: "",
  allowed_ips: "",
  mtu: "",
  dns: "",
  keepalive: "",
}

const fields: PeerField[] = [
  ["private_addr", "内网地址"],
  ["public_addr", "公网地址"],
  ["port", "监听端口"],
  ["allowed_ips", "默认路由"],
  ["mtu", "MTU"],
  ["dns", "DNS"],
  ["keepalive", "Keepalive"],
]
const IPV4_PATTERN = /^(\d{1,3}\.){3}\d{1,3}$/
const HOST_LABEL_PATTERN = /^[a-z0-9](?:[a-z0-9-]*[a-z0-9])?$/i
const MAX_PORT = 65535
const MIN_MTU = 64
const MAX_MTU = 1518
const MAX_HOST_LENGTH = 253

const isIPv4 = (value: string) =>
  IPV4_PATTERN.test(value) && value.split(".").every((part) => Number(part) <= 255)
const isHost = (value: string) =>
  value.length <= MAX_HOST_LENGTH &&
  value.split(".").every((label) => HOST_LABEL_PATTERN.test(label))

export default function Peers() {
  const [peers, setPeers] = useState<PeerItem[]>([])
  const [users, setUsers] = useState<UserItem[]>([])
  const [form, setForm] = useState<PeerFormState>(empty)
  const [editing, setEditing] = useState<PeerItem | null>(null)
  const [deleting, setDeleting] = useState<PeerItem | null>(null)
  const [formOpen, setFormOpen] = useState(false)

  const load = useCallback(async (): Promise<void> => {
    try {
      const raw = await (PeerListAPI as () => Promise<unknown>)()
      const response = raw as ApiResponse<PeerItem[]>
      if (response.status === 1) setPeers(response.data)
    } catch {
      toast.error("节点数据载入失败")
    }
  }, [])

  useEffect(() => {
    void load()
    const loadUsers = async () => {
      try {
        const raw = await (UserListAPI as () => Promise<unknown>)()
        const response = raw as ApiResponse<UserItem[]>
        if (response.status === 1) setUsers(response.data)
      } catch {
        toast.error("用户数据载入失败")
      }
    }
    void loadUsers()
  }, [load])

  const updateField = (key: PeerFieldKey | "user_id", value: string) => {
    setForm((current) => ({ ...current, [key]: value }))
  }

  const save = async () => {
    const userId = form.user_id.trim()
    const privateAddr = form.private_addr.trim()
    const publicAddr = form.public_addr.trim()
    const port = form.port.trim()
    const allowedIps = form.allowed_ips.trim()
    const mtu = form.mtu.trim()
    const dns = form.dns.trim()
    const keepalive = form.keepalive.trim()

    if (!editing && !userId) return toast.warning("请选择用户")
    if (!privateAddr) return toast.warning("请输入内网地址")
    if (!isIPv4(privateAddr)) return toast.warning("请输入正确的内网地址")
    if (publicAddr && !isHost(publicAddr)) return toast.warning("请输入正确的公网地址或域名")
    if (publicAddr && !port) return toast.warning("请输入公网地址和端口")
    if (Number(port) > MAX_PORT || Number(port) < 0) return toast.warning("请输入正确的端口")
    if (mtu && (Number(mtu) < MIN_MTU || Number(mtu) > MAX_MTU))
      return toast.warning("请输入正确的MTU")
    if (Number(keepalive) < 0) return toast.warning("请输入正确的Keepalive")

    const body = new models.Peers({
      id: editing?.id ?? 0,
      user_id: Number(userId),
      private_addr: privateAddr,
      public_addr: publicAddr,
      port: Number(port),
      allowed_ips: allowedIps,
      mtu: Number(mtu),
      dns,
      keepalive: Number(keepalive),
    })

    const request = editing
      ? (PeerUpdateAPI as (id: string, peer: models.Peers) => Promise<unknown>)(
          String(editing.id),
          body,
        )
      : (PeerAddAPI as (peer: models.Peers) => Promise<unknown>)(body)

    try {
      const raw = await request
      const response = raw as ApiResponse<null>
      if (response.status === 1) {
        setForm({ ...empty })
        setEditing(null)
        setFormOpen(false)
        void load()
        toast.success(response.message)
      } else toast.error(response.message)
    } catch {
      toast.error("节点接口请求失败")
    }
  }

  const remove = async () => {
    if (!deleting) return
    try {
      const raw = await (PeerDeleteAPI as (id: string) => Promise<unknown>)(String(deleting.id))
      const response = raw as ApiResponse<null>
      if (response.status === 1) {
        setDeleting(null)
        void load()
        toast.success(response.message)
      } else toast.error(response.message)
    } catch {
      toast.error("节点接口请求失败")
    }
  }

  const openEdit = (peer: PeerItem) => {
    setEditing(peer)
    setForm({
      user_id: "",
      username: peer.username,
      private_addr: peer.private_addr,
      public_addr: peer.public_addr,
      port: peer.public_addr ? String(peer.port) : "",
      allowed_ips: peer.allowed_ips,
      mtu: String(peer.mtu),
      dns: peer.dns,
      keepalive: String(peer.keepalive),
    })
    setFormOpen(true)
  }

  const formContent = (
    <div className="grid w-full gap-4 sm:grid-cols-2">
      {!editing && (
        <SelectField
          label="用户名"
          value={form.user_id}
          onChange={(value) => {
            updateField("user_id", value)
          }}
          options={users.map((user) => ({ value: String(user.id), label: user.username }))}
        />
      )}
      {editing && <Input label="用户名" value={form.username} disabled onChange={() => {}} />}
      {fields.map(([key, label]) => (
        <Input
          key={key}
          label={label}
          className="w-full"
          maxLength={
            key === "public_addr"
              ? MAX_HOST_LENGTH
              : key === "private_addr" || key === "allowed_ips" || key === "dns"
                ? 18
                : undefined
          }
          min={key === "mtu" ? 64 : 0}
          max={key === "port" ? MAX_PORT : key === "mtu" ? MAX_MTU : undefined}
          type={key === "port" || key === "mtu" || key === "keepalive" ? "number" : "text"}
          value={form[key] ?? ""}
          onChange={(event) => {
            updateField(key, event.target.value)
          }}
        />
      ))}
    </div>
  )

  return (
    <main className="mx-auto max-w-6xl px-5 py-6 sm:px-8">
      <header className="mb-7 flex flex-col gap-4 border-b border-slate-200 pb-5 sm:flex-row sm:items-end sm:justify-between dark:border-slate-700">
        <div>
          <p className="text-xs font-semibold uppercase tracking-[0.18em] text-wg-accent-strong">
            Network inventory
          </p>
          <h1 className="mt-1 text-2xl font-semibold tracking-tight text-slate-900 dark:text-white">
            节点
          </h1>
          <p className="mt-1 text-sm text-slate-500">管理本地接口与远程 WireGuard 节点</p>
        </div>
        <div className="flex items-center gap-3">
          <span className="ui-count">{peers.length} 个节点</span>
          <Button
            onClick={() => {
              setEditing(null)
              setForm(empty)
              setFormOpen(true)
            }}
          >
            增加节点
          </Button>
        </div>
      </header>
      <div className="grid gap-5 md:grid-cols-2 xl:grid-cols-3">
        {peers.map((peer) => (
          <Card key={peer.id} className="flex min-h-[330px] flex-col">
            <div className="flex items-start justify-between border-b border-slate-200 p-5 dark:border-slate-700">
              <div>
                <p className="text-xs font-medium uppercase tracking-wider text-slate-400">
                  Peer {peer.id}
                </p>
                <h2 className="mt-1 text-lg font-semibold text-slate-900 dark:text-white">
                  {peer.username}
                </h2>
              </div>
              <span
                className={`ui-status ${peer.public_addr ? "ui-status-active" : "ui-status-muted"}`}
              >
                {peer.public_addr ? "已配置公网" : "仅内网"}
              </span>
            </div>
            <div className="flex-1 space-y-4 p-5">
              <div className="grid grid-cols-2 gap-3">
                <div>
                  <p className="text-xs text-slate-400">内网地址</p>
                  <p className="mt-1 font-mono text-sm text-slate-800 dark:text-slate-200">
                    {peer.private_addr}
                  </p>
                </div>
                <div>
                  <p className="text-xs text-slate-400">公网端点</p>
                  <p className="mt-1 break-all font-mono text-sm text-slate-800 dark:text-slate-200">
                    {peer.public_addr ? `${peer.public_addr}:${peer.port}` : "未设置公网"}
                  </p>
                </div>
              </div>
              <div className="ui-inset">
                <p className="text-xs text-slate-400">默认路由</p>
                <p className="mt-1 break-all font-mono text-sm text-slate-700 dark:text-slate-300">
                  {peer.allowed_ips || "未设置"}
                </p>
              </div>
              <div className="grid grid-cols-3 gap-2 text-center text-xs">
                <div className="ui-inset p-2">
                  <p className="text-slate-400">DNS</p>
                  <p className="mt-1 break-all text-slate-700 dark:text-slate-300">
                    {peer.dns || "未设置"}
                  </p>
                </div>
                <div className="ui-inset p-2">
                  <p className="text-slate-400">MTU</p>
                  <p className="mt-1 text-slate-700 dark:text-slate-300">{peer.mtu || "未设置"}</p>
                </div>
                <div className="ui-inset p-2">
                  <p className="text-slate-400">Keepalive</p>
                  <p className="mt-1 text-slate-700 dark:text-slate-300">{peer.keepalive}</p>
                </div>
              </div>
            </div>
            <div className="ui-action-bar">
              <Button
                variant="outline"
                onClick={() => {
                  openEdit(peer)
                }}
              >
                修改
              </Button>
              <Button
                variant="danger"
                onClick={() => {
                  setDeleting(peer)
                }}
              >
                删除
              </Button>
            </div>
          </Card>
        ))}
      </div>
      <Dialog
        open={formOpen}
        onClose={() => {
          setEditing(null)
          setForm(empty)
          setFormOpen(false)
        }}
        title={editing ? "修改节点" : "增加节点"}
        className="max-w-2xl"
      >
        {formContent}
        <div className="ui-dialog-actions mt-5">
          <Button
            variant="outline"
            onClick={() => {
              setEditing(null)
              setForm(empty)
              setFormOpen(false)
            }}
          >
            取消
          </Button>
          <Button
            onClick={() => {
              void save()
            }}
          >
            提交
          </Button>
        </div>
      </Dialog>
      <Dialog
        open={deleting !== null}
        onClose={() => {
          setDeleting(null)
        }}
        title="确认删除节点"
      >
        <div className="ui-dialog-actions">
          <Button
            variant="outline"
            onClick={() => {
              setDeleting(null)
            }}
          >
            取消
          </Button>
          <Button variant="danger" onClick={() => void remove()}>
            提交
          </Button>
        </div>
      </Dialog>
    </main>
  )
}
