import { zodResolver } from "@hookform/resolvers/zod"
import axios from "axios"
import { useState } from "react"
import { useFieldArray, useForm } from "react-hook-form"
import { z } from "zod"
import styles from "./App.module.css"
import { Result, Room } from "./types"

const schema = z.object({
  room: Room,
})
type formSchema = z.infer<typeof schema>

const inputOptions = {
  valueAsNumber: true,
}

function App() {
  const {
    register,
    control,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<formSchema>({
    defaultValues: {
      room: [...Array(4).keys()].map(() => ({ width: 1, height: 1, qtyWindows: 0, qtyDoors: 0 })),
    },
    resolver: zodResolver(schema),
  })
  const { fields } = useFieldArray({ control, name: "room" })

  const [results, setResults] = useState<Result[]>([])

  return (
    <div className={styles.container}>
      <div>
        <h1>Calculadora de necessidades de tinta</h1>
      </div>
      <form
        className={styles.form}
        onSubmit={handleSubmit(async (data) => {
          const res = await axios.post(`${import.meta.env.VITE_API_URL}/cans-needed`, data)

          setResults((r) => [res.data, ...r.slice(0, 5)])
        })}
      >
        <table id="form">
          <thead>
            <tr>
              <th>Parede</th>
              <th>Largura (m)</th>
              <th>Altura (m)</th>
              <th>Qtd. de janelas (unid.)</th>
              <th>Qtd. de portas (unid.)</th>
            </tr>
          </thead>
          <tbody>
            {fields.map((item, index) => (
              <tr key={item.id}>
                <td>Nº {index + 1}</td>
                <td>
                  <input type="text" autoComplete="off" {...register(`room.${index}.width`, inputOptions)} />
                  <Error message={errors.room?.[index]?.width?.message} />
                </td>
                <td>
                  <input type="text" autoComplete="off" {...register(`room.${index}.height`, inputOptions)} />
                  <Error message={errors.room?.[index]?.height?.message} />
                </td>
                <td>
                  <input
                    type="text"
                    autoComplete="off"
                    {...register(`room.${index}.qtyWindows`, inputOptions)}
                  />
                  <Error message={errors.room?.[index]?.qtyWindows?.message} />
                </td>
                <td>
                  <input
                    type="text"
                    autoComplete="off"
                    {...register(`room.${index}.qtyDoors`, inputOptions)}
                  />
                  <Error message={errors.room?.[index]?.qtyDoors?.message} />
                </td>
              </tr>
            ))}
          </tbody>
        </table>

        <div className={styles.submit}>
          <button type="button" disabled={results.length < 1} onClick={() => setResults([])}>
            Limpar resultados
          </button>
          <button type="submit">Calcular</button>
          <button type="button" onClick={() => reset()}>
            Reset
          </button>
        </div>
      </form>
      <Results results={results} />
    </div>
  )
}

function Error({ message }: { message: string | undefined }) {
  return message ? <div className={styles.error}>{message}</div> : null
}

function Results({ results }: { results: Result[] }) {
  return (
    <div className={styles.results}>
      {results.map((r) => (
        <Result key={r.id} result={r} />
      ))}
    </div>
  )
}

function Result({ result }: { result: Result }) {
  return (
    <div className={styles.result}>
      <div>
        <p>Necessidade</p>
        <p>{result.target}L</p>
      </div>
      <div>
        <p>Solução</p>
        <p>{result.solved_for}L</p>
      </div>
      <table id="result">
        <thead>
          <tr>
            <th>Lata</th>
            <th>Qtd.</th>
          </tr>
        </thead>
        <tbody>
          {result.cans.map((c) => (
            <tr key={c.id}>
              <td>{c.label}</td>
              <td>{c.qty}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}

export default App
