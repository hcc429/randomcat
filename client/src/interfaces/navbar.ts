

export interface NavItem{
    to: string
    text: string
    onclick?: (...args: any[])=>any
}