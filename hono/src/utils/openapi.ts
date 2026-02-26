import { describeRoute, resolver } from "hono-openapi";

export function createRouteDescriber(tags: string[], successSchema: unknown) {
  return (summary: string) =>
    describeRoute({
      tags,
      summary,
      responses: {
        200: {
          description: "Successful response",
          content: {
            "application/json": {
              schema: resolver(successSchema as never)
            }
          }
        }
      }
    });
}
