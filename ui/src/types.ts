import { z } from "zod"

export const Wall = z.object({
  width: z.number().int().positive(),
  height: z.number().int().positive(),
  qtyWindows: z.number().int().nonnegative(),
  qtyDoors: z.number().int().nonnegative(),
})
export type Wall = z.infer<typeof Wall>

export const Room = z.array(Wall).length(4)
export type Room = z.infer<typeof Room>

export type Result = {
  id: string
  target: number
  solved_for: number
  cans: CanLine[]
}

export type CanLine = {
  id: string
  label: string
  qty: number
}
