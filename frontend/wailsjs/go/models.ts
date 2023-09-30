export namespace calc {
	
	export class Token {
	    kind: string;
	    val: number;
	    str: string;
	
	    static createFrom(source: any = {}) {
	        return new Token(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.kind = source["kind"];
	        this.val = source["val"];
	        this.str = source["str"];
	    }
	}

}

export namespace dev {
	
	export class Node {
	    tokens: calc.Token[];
	    buffer: string;
	    result: number;
	    name: string;
	    operator: string;
	    done: boolean;
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new Node(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.tokens = this.convertValues(source["tokens"], calc.Token);
	        this.buffer = source["buffer"];
	        this.result = source["result"];
	        this.name = source["name"];
	        this.operator = source["operator"];
	        this.done = source["done"];
	        this.message = source["message"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class AppState {
	    nodes: Node[];
	    nodeIndex: number;
	    value: number;
	    bufferIndex: number;
	    clearing: boolean;
	    hideName: boolean;
	
	    static createFrom(source: any = {}) {
	        return new AppState(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.nodes = this.convertValues(source["nodes"], Node);
	        this.nodeIndex = source["nodeIndex"];
	        this.value = source["value"];
	        this.bufferIndex = source["bufferIndex"];
	        this.clearing = source["clearing"];
	        this.hideName = source["hideName"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

