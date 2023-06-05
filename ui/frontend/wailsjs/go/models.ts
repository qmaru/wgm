export namespace models {
	
	export class Peers {
	    id: number;
	    created_at: number;
	    updated_at: number;
	    remark: string;
	    state: number;
	    user_id: number;
	    public_addr: string;
	    private_addr: string;
	    port: number;
	    allowed_ips: string;
	    mtu: number;
	    dns: string;
	    keepalive: number;
	
	    static createFrom(source: any = {}) {
	        return new Peers(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	        this.remark = source["remark"];
	        this.state = source["state"];
	        this.user_id = source["user_id"];
	        this.public_addr = source["public_addr"];
	        this.private_addr = source["private_addr"];
	        this.port = source["port"];
	        this.allowed_ips = source["allowed_ips"];
	        this.mtu = source["mtu"];
	        this.dns = source["dns"];
	        this.keepalive = source["keepalive"];
	    }
	}
	export class Routes {
	    id: number;
	    created_at: number;
	    updated_at: number;
	    remark: string;
	    state: number;
	    cidr: string;
	
	    static createFrom(source: any = {}) {
	        return new Routes(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	        this.remark = source["remark"];
	        this.state = source["state"];
	        this.cidr = source["cidr"];
	    }
	}
	export class Users {
	    id: number;
	    created_at: number;
	    updated_at: number;
	    remark: string;
	    state: number;
	    username: string;
	    prikey: string;
	    pubkey: string;
	
	    static createFrom(source: any = {}) {
	        return new Users(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	        this.remark = source["remark"];
	        this.state = source["state"];
	        this.username = source["username"];
	        this.prikey = source["prikey"];
	        this.pubkey = source["pubkey"];
	    }
	}

}

