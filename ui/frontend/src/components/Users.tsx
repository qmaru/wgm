import { useCallback, useEffect, useState } from "react"
import { toast } from "sonner"

import { UserAddAPI, UserDeleteAPI, UserListAPI, UserUpdateAPI } from "@wails/go/backend/App"
import { models } from "@wails/go/models"

import { ApiResponse, UserItem } from "@/types"
import { Button, Card, Dialog, Input } from "@/components/ui"

export default function Users() {
  const [users, setUsers] = useState<UserItem[]>([])
  const [username, setUsername] = useState("")
  const [editing, setEditing] = useState<UserItem | null>(null)
  const [deleting, setDeleting] = useState<UserItem | null>(null)

  const load = useCallback(async (): Promise<void> => {
    try {
      const raw = await (UserListAPI as () => Promise<unknown>)()
      const response = raw as ApiResponse<UserItem[]>
      if (response.status === 1) setUsers(response.data)
    } catch {
      toast.error("用户接口请求失败")
    }
  }, [])

  useEffect(() => {
    void load()
  }, [load])

  const add = async () => {
    const value = username.trim()
    if (!value) return toast.warning("请输入用户名")
    const body = new models.Users({ username: value })
    try {
      const raw = await (UserAddAPI as (user: models.Users) => Promise<unknown>)(body)
      const response = raw as ApiResponse<null>
      if (response.status === 1) {
        setUsername("")
        void load()
        toast.success(response.message)
      } else toast.error(response.message)
    } catch {
      toast.error("用户接口请求失败")
    }
  }

  const update = async () => {
    if (!editing) return
    const value = editing.username.trim()
    if (!value) return toast.warning("请输入用户名")
    const body = new models.Users({ id: editing.id, username: value })
    try {
      const raw = await (UserUpdateAPI as (id: string, user: models.Users) => Promise<unknown>)(
        String(editing.id),
        body,
      )
      const response = raw as ApiResponse<null>
      if (response.status === 1) {
        setEditing(null)
        void load()
        toast.success(response.message)
      } else toast.error(response.message)
    } catch {
      toast.error("用户接口请求失败")
    }
  }

  const copyKey = async (key: string): Promise<void> => {
    try {
      await navigator.clipboard.writeText(key)
      toast.success("密钥已复制")
    } catch {
      toast.error("密钥复制失败")
    }
  }

  const displayKey = (key: string) =>
    key.length > 10 ? `${key.slice(0, 4)}...${key.slice(-6)}` : key

  const remove = async () => {
    if (!deleting) return
    try {
      const raw = await (UserDeleteAPI as (id: string) => Promise<unknown>)(String(deleting.id))
      const response = raw as ApiResponse<null>
      if (response.status === 1) {
        setDeleting(null)
        void load()
        toast.success(response.message)
      } else toast.error(response.message)
    } catch {
      toast.error("用户接口请求失败")
    }
  }

  return (
    <main className="mx-auto max-w-6xl px-5 py-6 sm:px-8">
      <header className="mb-7 flex flex-col gap-4 border-b border-slate-200 pb-5 sm:flex-row sm:items-end sm:justify-between dark:border-slate-700">
        <div>
          <p className="text-xs font-semibold uppercase tracking-[0.18em] text-wg-accent-strong">
            Identity vault
          </p>
          <h1 className="mt-1 text-2xl font-semibold tracking-tight text-slate-900 dark:text-white">
            用户
          </h1>
          <p className="mt-1 text-sm text-slate-500">管理用户身份与 WireGuard 密钥</p>
        </div>
        <span className="ui-count self-start">{users.length} 个用户</span>
      </header>
      <div className="mb-8 w-full xl:w-[calc((100%-3.75rem)/4)]">
        <div className="ui-input-group">
          <Input
            label="用户名"
            className="w-full"
            value={username}
            onChange={(event) => {
              setUsername(event.target.value)
            }}
          />
          <Button
            className="ui-input-action whitespace-nowrap"
            onClick={() => {
              void add()
            }}
          >
            提交
          </Button>
        </div>
      </div>
      <div className="grid gap-5 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
        {users.map((user) => (
          <Card key={user.id} className="flex min-h-[285px] flex-col">
            <div className="flex-1 p-5">
              <div>
                <p className="text-xs text-slate-400">USER {user.id}</p>
                <h2 className="mt-1 text-lg font-semibold text-slate-900 dark:text-white">
                  {user.username}
                </h2>
              </div>
              <div className="mt-6 space-y-4 text-xs">
                <div>
                  <p className="mb-1 font-medium text-wg-accent-strong">Private key</p>
                  <button
                    type="button"
                    title="点击复制完整 Private key"
                    onClick={() => {
                      void copyKey(user.private_key)
                    }}
                    className="ui-preview ui-copyable block w-full break-all p-3 text-left font-mono leading-5 text-slate-600 dark:text-slate-300"
                  >
                    {displayKey(user.private_key)}
                  </button>
                </div>
                <div>
                  <p className="mb-1 font-medium text-wg-accent-strong">Public key</p>
                  <button
                    type="button"
                    title="点击复制完整 Public key"
                    onClick={() => {
                      void copyKey(user.public_key)
                    }}
                    className="ui-preview ui-copyable block w-full break-all p-3 text-left font-mono leading-5 text-slate-600 dark:text-slate-300"
                  >
                    {displayKey(user.public_key)}
                  </button>
                </div>
              </div>
            </div>
            <div className="ui-action-bar">
              <Button
                variant="outline"
                onClick={() => {
                  setEditing(user)
                }}
              >
                修改
              </Button>
              <Button
                variant="danger"
                onClick={() => {
                  setDeleting(user)
                }}
              >
                删除
              </Button>
            </div>
          </Card>
        ))}
      </div>
      <Dialog
        open={editing !== null}
        onClose={() => {
          setEditing(null)
        }}
        title="修改用户"
        className="max-w-sm"
      >
        {editing && (
          <>
            <Input
              label="用户名"
              className="my-2 w-64"
              value={editing.username}
              onChange={(event) => {
                setEditing({ ...editing, username: event.target.value })
              }}
            />
            <div className="ui-dialog-actions mt-5">
              <Button
                variant="outline"
                onClick={() => {
                  setEditing(null)
                }}
              >
                取消
              </Button>
              <Button
                onClick={() => {
                  void update()
                }}
              >
                提交
              </Button>
            </div>
          </>
        )}
      </Dialog>
      <Dialog
        open={deleting !== null}
        onClose={() => {
          setDeleting(null)
        }}
        title="确认删除用户"
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
