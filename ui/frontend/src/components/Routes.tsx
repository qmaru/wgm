import { useCallback, useEffect, useState } from "react"
import { toast } from "sonner"

import { RouteAddAPI, RouteDeleteAPI, RouteListAPI, RouteUpdateAPI } from "@wails/go/backend/App"
import { models } from "@wails/go/models"

import { ApiResponse, RouteItem } from "@/types"
import { Button, Card, Dialog, Input } from "@/components/ui"

export default function Routes() {
  const [routes, setRoutes] = useState<RouteItem[]>([])
  const [cidr, setCidr] = useState("")
  const [editing, setEditing] = useState<RouteItem | null>(null)
  const [deleting, setDeleting] = useState<RouteItem | null>(null)

  const load = useCallback(async (): Promise<void> => {
    try {
      const raw = await (RouteListAPI as () => Promise<unknown>)()
      const response = raw as ApiResponse<RouteItem[]>
      if (response.status === 1) setRoutes(response.data)
    } catch {
      toast.error("路由接口请求失败")
    }
  }, [])

  useEffect(() => {
    void load()
  }, [load])

  const add = async () => {
    const value = cidr.trim()
    if (!value) return toast.warning("请输入路由")
    const body = new models.Routes({ cidr: value })
    try {
      const raw = await (RouteAddAPI as (route: models.Routes) => Promise<unknown>)(body)
      const response = raw as ApiResponse<null>
      if (response.status === 1) {
        setCidr("")
        void load()
        toast.success(response.message)
      } else toast.error(response.message)
    } catch {
      toast.error("路由接口请求失败")
    }
  }

  const update = async () => {
    if (!editing) return
    const value = editing.cidr.trim()
    if (!value) return toast.warning("请输入路由")
    const body = new models.Routes({ id: editing.id, cidr: value })
    try {
      const raw = await (RouteUpdateAPI as (id: string, route: models.Routes) => Promise<unknown>)(
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
      toast.error("路由接口请求失败")
    }
  }

  const remove = async () => {
    if (!deleting) return
    try {
      const raw = await (RouteDeleteAPI as (id: string) => Promise<unknown>)(String(deleting.id))
      const response = raw as ApiResponse<null>
      if (response.status === 1) {
        setDeleting(null)
        void load()
        toast.success(response.message)
      } else toast.error(response.message)
    } catch {
      toast.error("路由接口请求失败")
    }
  }

  return (
    <main className="mx-auto max-w-6xl px-5 py-6 sm:px-8">
      <header className="mb-7 flex flex-col gap-4 border-b border-slate-200 pb-5 sm:flex-row sm:items-end sm:justify-between dark:border-slate-700">
        <div>
          <p className="text-xs font-semibold uppercase tracking-[0.18em] text-wg-accent-strong">
            Routing table
          </p>
          <h1 className="mt-1 text-2xl font-semibold tracking-tight text-slate-900 dark:text-white">
            路由
          </h1>
          <p className="mt-1 text-sm text-slate-500">维护可分发到节点的 CIDR 路由规则</p>
        </div>
        <span className="ui-count self-start">{routes.length} 条规则</span>
      </header>
      <div className="mb-8 w-full xl:w-[calc((100%-3.75rem)/4)] border-b border-dashed border-slate-300 pb-6 dark:border-slate-700">
        <div className="ui-input-group">
          <Input
            label="路由"
            className="w-full"
            maxLength={18}
            value={cidr}
            onChange={(event) => {
              setCidr(event.target.value)
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
        {routes.map((route) => (
          <Card
            key={route.id}
            className="flex min-h-[190px] flex-col border-l-4 border-l-wg-accent"
          >
            <div className="flex-1 p-5">
              <span className="font-mono text-xs font-semibold uppercase tracking-wider text-wg-accent-strong">
                CIDR rule
              </span>
              <p className="mt-7 break-all font-mono text-xl font-semibold text-slate-900 dark:text-white">
                {route.cidr}
              </p>
              {route.remark && <p className="mt-2 text-sm text-slate-500">{route.remark}</p>}
            </div>
            <div className="ui-action-bar">
              <Button
                variant="outline"
                onClick={() => {
                  setEditing(route)
                }}
              >
                修改
              </Button>
              <Button
                variant="danger"
                onClick={() => {
                  setDeleting(route)
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
        title="修改路由"
        className="max-w-sm"
      >
        {editing && (
          <>
            <Input
              label="路由"
              className="my-2 w-64"
              maxLength={18}
              value={editing.cidr}
              onChange={(event) => {
                setEditing({ ...editing, cidr: event.target.value })
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
        title="确认删除路由"
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
