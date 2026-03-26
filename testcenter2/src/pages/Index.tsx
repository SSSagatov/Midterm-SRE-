import { useState } from "react";
import utoLogo from "@/pages/20260315_1944_Image Generation_remix_01kkrz7w0kfq78bqz2csx3r367.png";
import pdf from "@/pages/1-2026-2-004234422-1-1.pdf"

const certificateData = {
  fullName: "КЕНЕБАЕВ РАСУЛ ДАНИЯРОВИЧ",
  iin: "090312553541",
  ikt: "004300689",
  testType: "ЕНТ",
  testYear: "2026 (Март)",
  language: "русский",
  validUntil: "31.12.2026",
};

const subjects = [
  { id: 1, name: "История Казахстана", score: 14 },
  { id: 2, name: "Грамотность чтения", score: 8 },
  { id: 3, name: "Математическая грамотность", score: 10 },
  { id: 4, name: "Всемирная история", score: 30 },
  { id: 5, name: "Основы права", score: 46 },
];

const totalScore = subjects.reduce((sum, s) => sum + s.score, 0);

const Index = () => {
  const [searchType, setSearchType] = useState<"ikt" | "serial">("ikt");

  return (
    <div className="min-h-svh bg-background">
      {/* Top Header Bar */}
      <header className="bg-primary bg-blue-500 text-primary-foreground">
        <div className="container mx-auto px-3 sm:px-4 md:px-8 lg:px-60 flex items-center justify-between h-12 sm:h-10">
          <div className="flex items-center gap-4 sm:gap-8 min-w-0">
            <div className="flex items-center gap-2 min-w-0">
              <img src={utoLogo} alt="UTO logo" className="h-8 sm:h-[50px] w-auto shrink-0" />

              <span className="text-[11px] leading-tight sm:hidden">
                Национальный центр тестирования
              </span>

              <span className="hidden sm:inline text-sm">
                Национальный центр тестирования
              </span>
            </div>

            <span className="text-xs hidden md:inline">
              Тел: <a href="tel:87172695069" className="hover:underline">8 (7172) 69-50-69</a>
            </span>

            <span className="text-xs hidden md:inline">
              ✉ <a href="mailto:uto@testcenter.kz" className="hover:underline">uto@testcenter.kz</a>
            </span>
          </div>

          <div className="flex items-center gap-3 shrink-0">
            <div className="relative ml-2">
              <select
                defaultValue="ru"
                className="
                  appearance-none
                  bg-[hsl(198_52.2%_36.6%)]
                  border-0
                  outline-none
                  text-xs
                  text-white
                  pr-5
                  pl-2
                  py-1
                  rounded
                  cursor-pointer
                  focus:outline-none
                "
              >
                <option value="ru">ru</option>
                <option value="kk">kz</option>
              </select>

              <span className="pointer-events-none absolute right-2 top-1/2 -translate-y-1/2 text-[10px] text-white">
                ▾
              </span>
            </div>
          </div>
        </div>
      </header>

      {/* Main Content */}
      <div className="container mx-auto px-1 py-6">
        <div className="flex flex-col lg:flex-row gap-6">
          {/* Left Sidebar - Search */}
          <aside className="w-full lg:w-80 shrink-0">
            <div className="mb-3">
              <label className="flex items-center gap-2 text-sm cursor-pointer mb-1">
                <input
                  type="radio"
                  name="searchType"
                  checked={searchType === "ikt"}
                  onChange={() => setSearchType("ikt")}
                  className="accent-primary"
                />
                поиск по ИКТ
              </label>
              <label className="flex items-center gap-2 text-sm cursor-pointer">
                <input
                  type="radio"
                  name="searchType"
                  checked={searchType === "serial"}
                  onChange={() => setSearchType("serial")}
                  className="accent-primary"
                />
                поиск по серии и по номеру сертификата
              </label>
            </div>

            <hr className="border-border mb-4" />

            <div>
              <h4 className="text-[22px] font-bold mb-2">поиск по ИКТ</h4>

              <div className="space-y-6">
                <div className="flex flex-col sm:flex-row sm:items-center gap-1">
                  <label className="text-sm w-40 shrink-0">Выберите тип тестирования</label>
                  <select className="border border-input rounded px-2 py-1.5 text-sm w-full bg-background">
                    <option>ЕНТ</option>
                    <option>Магистратура/Докторантура</option>
                    <option>ОЗП</option>
                    <option>QAZTEST</option>
                    <option>Кандидаты</option>
                    <option>Гражданство</option>
                  </select>
                </div>

                <div className="flex flex-col sm:flex-row sm:items-center gap-1">
                  <label className="text-sm w-40 shrink-0">Год</label>
                  <select className="border border-input rounded px-2 py-1.5 text-sm w-full bg-background">
                    {Array.from({ length: 14 }, (_, i) => 2026 - i).map((year) => (
                      <option key={year}>{year}</option>
                    ))}
                  </select>
                </div>

                <div className="flex flex-col sm:flex-row sm:items-center gap-1">
                  <label className="text-sm w-40 shrink-0">Введите ИИН</label>
                  <input
                    type="text"
                    placeholder="ИИН"
                    maxLength={12}
                    className="border border-input rounded px-2 py-1.5 text-sm w-full bg-background"
                  />
                </div>

                <div className="flex flex-col sm:flex-row sm:items-center gap-1">
                  <label className="text-sm w-40 shrink-0">Введите ИКТ</label>
                  <input
                    type="text"
                    placeholder="ИКТ"
                    maxLength={9}
                    className="border border-input rounded px-2 py-1.5 text-sm w-full bg-background"
                  />
                </div>

                <button className="border border-primary text-primary rounded px-4 py-1.5 text-sm hover:bg-primary hover:text-primary-foreground transition-colors">
                  Поиск
                </button>
              </div>
            </div>
          </aside>

          {/* Center - Certificate Data & Subjects side by side */}
          <div className="flex-1 min-w-0">
            <div className="flex flex-col lg:flex-row gap-8">
              {/* Certificate Data - plain text with underlines */}
              <div className="shrink-0" style={{ width: 'fit-content' }}>
                <h5 className="text-lg font-bold mb-2">Данные сертификата</h5>
                <div className="space-y-0">
                  {[
                    ["ФИО", certificateData.fullName],
                    ["ИИН", certificateData.iin],
                    ["ИКТ", certificateData.ikt],
                    ["Вид тестирования", certificateData.testType],
                    ["Год тестирования", certificateData.testYear],
                    ["Язык сдачи тестирования", certificateData.language],
                    ["Действителен до", certificateData.validUntil],
                  ].map(([label, value], i) => (
                    <div key={i} className="border-b border-border py-2 text-sm">
                      <span className="font-bold">{label}</span>: {value}
                    </div>
                  ))}
                </div>
              </div>

              {/* Subjects Table */}
              <div className="w-full lg:w-80 shrink-0">
                <h5 className="text-lg font-bold mb-4">Предметы тестирования</h5>
                <table className="w-full border border-border text-sm">
                  <thead>
                    <tr className="border-b border-border">
                      <td className="border-r border-border px-2 py-1.5 font-bold text-center w-10">№</td>
                      <td className="border-r border-border px-2 py-1.5 font-bold">Название предмета</td>
                      <td className="px-2 py-1.5 font-bold text-center w-14">Балл</td>
                    </tr>
                  </thead>
                  <tbody>
                    {subjects.map((s) => (
                      <tr key={s.id} className="border-b border-border">
                        <td className="border-r border-border px-2 py-1.5 text-center">{s.id}</td>
                        <td className="border-r border-border px-2 py-1.5">{s.name}</td>
                        <td className="px-2 py-1.5 text-center">{s.score}</td>
                      </tr>
                    ))}
                    <tr>
                      <td colSpan={2} className="border-r border-border px-2 py-1.5 text-right">
                        <span className="font-bold">Общий балл по предметам:</span>
                      </td>
                      <td className="px-2 py-1.5 text-center font-bold">{totalScore}</td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>

            <a
              href="https://1-2026-2-004300689-1-1-pdf.pages.dev/1-2026-2-004300689-1-1.pdf"
              target="_blank"
              rel="noopener noreferrer"
              className="inline-block mt-4 border border-primary text-primary rounded px-5 py-2 text-sm hover:bg-primary hover:text-primary-foreground transition-colors"
            >
              Скачать документ
            </a>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Index;
