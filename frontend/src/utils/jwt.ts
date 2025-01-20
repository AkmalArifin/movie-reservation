export function parseJwt(token: string | null): {
    id: number,
    role: string,
    iat: number,
    exp: number
} {
    if (token == null) {
        return {
            id: 0,
            role: '',
            iat: 0,
            exp: 0
        }
    } else {
        return JSON.parse(atob(token.split('.')[1]))
    }
}