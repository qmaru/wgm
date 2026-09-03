export type ApiResponse<T> = {
  status: number
  message: string
  data: T
}

export type UserItem = {
  id: number
  username: string
  private_key: string
  public_key: string
}

export type RouteItem = {
  id: number
  cidr: string
  remark?: string
}

export type PeerItem = {
  id: number
  username: string
  prikey: string
  pubkey: string
  private_addr: string
  public_addr: string
  port: number
  allowed_ips: string
  mtu: number
  dns: string
  keepalive: number
}

export type DataList = {
  peers: PeerItem[]
  routes: RouteItem[]
}

export type PeerFieldKey =
  | "private_addr"
  | "public_addr"
  | "port"
  | "allowed_ips"
  | "mtu"
  | "dns"
  | "keepalive"

export type PeerFormState = {
  user_id: string
  username: string
  private_addr: string
  public_addr: string
  port: string
  allowed_ips: string
  mtu: string
  dns: string
  keepalive: string
}

export type PeerField = [PeerFieldKey, string]
